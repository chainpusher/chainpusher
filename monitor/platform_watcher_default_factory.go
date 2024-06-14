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
	if p.Config.IsTesting { // this demo service is used for testing, only one transaction is processed
		transactionService = application.NewTransactionDemoService(p.Config)
	} else { // it's normal service
		transactionService = application.NewDefaultTransactionService(p.Config)
	}

	switch platform {
	case model.PlatformTron:
		return p.CreteTronPlatformWatcher(platform, transactionService)
	case model.PlatformEthereum:
		return p.CreteEthereumPlatformWatcher(platform, transactionService)
	default:
		panic("Platform not supported")
	}

}

func (p *PlatformWatcherDefaultFactory) CreteTronPlatformWatcher(
	platform model.Platform,
	transactionService application.TransactionService,
) PlatformWatcher {

	w := NewWatcher(NewWatcherOptions{
		TimeRequiredForOneCycle: 3,
		IsOneTime:               false,
		TickHandler:             nil,
	})
	tronBlockchainService := chain.NewTronV2BlockChainService()
	return NewPlatformWatcherV2(w, tronBlockchainService, transactionService)
}

func NewPlatformWatcherDefaultFactory(ctx *Ctx) *PlatformWatcherDefaultFactory {
	return &PlatformWatcherDefaultFactory{
		Config:  ctx.Config,
		Channel: ctx.Channel,
	}
}

func (p *PlatformWatcherDefaultFactory) CreteEthereumPlatformWatcher(
	platform model.Platform,
	transactionService application.TransactionService,
) PlatformWatcher {
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
}
