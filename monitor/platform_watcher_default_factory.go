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
	Config  *config.Config
	Channel chan interface{}
}

func (p *PlatformWatcherDefaultFactory) CreatePlatformWatcher(platform model.Platform) PlatformWatcher {

	var transactionService application.TransactionService
	if p.Config.IsTesting {
		transactionService = application.NewTransactionDemoService(p.Config)
	} else {
		transactionService = application.NewDefaultTransactionService(p.Config)
	}

	switch platform {
	case model.PlatformTron:
		return NewPlatformWatcherTron(p.Config, transactionService, p.Channel, p.Config.IsTesting)
	case model.PlatformEthereum:
		infuraApiUrl := chain.GetInfuraApiUrlV2(p.Config.InfuraKey)
		ethereumService, err := chain.NewEthereumBlockChainService(infuraApiUrl, p.Channel)
		if err != nil {
			panic(fmt.Sprintf("Error creating Ethereum service: %v", err))
		}

		return NewPlatformWatcherEthereum(
			15*time.Second,
			nil,
			ethereumService,
			transactionService,
			p.Config.IsTesting,
		)
	default:
		panic("Platform not supported")
	}

}

func NewPlatformWatcherDefaultFactory(ctx *Ctx) *PlatformWatcherDefaultFactory {
	return &PlatformWatcherDefaultFactory{
		Config:  ctx.Config,
		Channel: ctx.Channel,
	}
}
