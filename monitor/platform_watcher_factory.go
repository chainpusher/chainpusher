package monitor

import (
	"errors"
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/application"
	"math/big"
	"time"
)

func NewPlatformWatcher(ctx *Ctx, platform model.Platform) (*PlatformWatcher, error) {
	var blockChainService service.BlockChainService
	var err error
	var transactionService application.AnalysisService
	var timeForBlockGenerated time.Duration

	if ctx.Config.IsTesting {
		transactionService = application.NewTransactionDemoService(ctx.Config)
	} else {
		transactionService = application.NewDefaultAnalysisService(ctx.Config)
	}

	switch platform {
	case model.PlatformEthereum:
		blockChainService, err = NewEthereumBlockChainService(ctx, ctx.Listeners)
		timeForBlockGenerated = 15
	case model.PlatformTron:
		blockChainService, err = NewTRONBlockChainService(ctx, ctx.Listeners)
		timeForBlockGenerated = 3
	default:
		blockChainService, err = nil, errors.New("platform not supported")
	}

	if err != nil {
		return nil, err
	}

	return &PlatformWatcher{
		config:                ctx.Config,
		done:                  make(chan bool),
		service:               blockChainService,
		applicationService:    transactionService,
		number:                big.NewInt(-1),
		once:                  ctx.Config.IsTesting,
		isRestart:             false,
		timeForBlockGenerated: timeForBlockGenerated,
		platform:              platform,
	}, nil
}

func NewEthereumBlockChainService(ctx *Ctx, listeners []service.BlockListener) (service.BlockChainService, error) {
	url := service.GetInfuraApiUrlV2(ctx.Config.InfuraKey)
	return service.NewEthereumBlockChainService(url, listeners)
}

func NewTRONBlockChainService(ctx *Ctx, listeners []service.BlockListener) (service.BlockChainService, error) {
	client, err := service.NewTronClient()
	if err != nil {
		return nil, err
	}
	return service.NewTronBlockChainService(client, listeners), nil
}
