package monitor

import (
	"fmt"
	"time"

	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/chain"
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/model"
)

type PlatformWatcherDefaultFactory struct {
	Config *config.Config
}

func (p *PlatformWatcherDefaultFactory) CreatePlatformWatcher(platform model.Platform) PlatformWatcher {
	transactionService := application.NewTransactionService(p.Config)

	infuraApiUrl := chain.GetInfuraApiUrlV2(p.Config.InfuraKey)
	ethereumService, err := chain.NewEthereumBlockChainService(infuraApiUrl)
	if err != nil {
		panic(fmt.Sprintf("Error creating Ethereum service: %v", err))
	}

	switch platform {
	case model.PlatformTron:
		return NewPlatformWatcherTron(p.Config, transactionService)
	case model.PlatformEthereum:
		return NewPlatformWatcherEthereum(
			15*time.Second,
			nil,
			ethereumService,
			transactionService,
		)
	default:
		panic("Platform not supported")
	}

}

func NewPlatformWatcherDefaultFactory(config *config.Config) *PlatformWatcherDefaultFactory {
	return &PlatformWatcherDefaultFactory{
		Config: config,
	}
}
