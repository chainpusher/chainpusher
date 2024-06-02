package monitor

import (
	"log"
	"time"

	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/chain"
	"github.com/chainpusher/chainpusher/config"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type PlatformWatcherTron struct {
	Config             *config.Config
	done               chan bool
	Service            *chain.TronBlockChainService
	ApplicationService *application.TransactionService
	Number             int64
}

func (p *PlatformWatcherTron) Start() {
	for {
		latest, transactions, err := p.Service.GetNowBlock()
		if err != nil {
			log.Printf("Error getting block: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("Latest block number: %d and transactions %d", latest.BlockHeader.RawData.Number, len(transactions))
		p.ApplicationService.AnalyzeTrades(transactions)
		p.Number = latest.BlockHeader.RawData.Number + 1
		break
	}

	p.WatchBlocks()

	log.Println("Starting Tron platform watcher ...")
}

func (p *PlatformWatcherTron) FetchBlocks() {
	logrus.Debug("Fetching block ", p.Number)
	transactions, err := p.Service.GetBlock(p.Number)
	if err != nil {
		logrus.Warnf("Error getting block: %v", err)
		return
	}
	logrus.Infof("Block %d fetched with %d transactions", p.Number, len(transactions))
	p.ApplicationService.AnalyzeTrades(transactions)
}

func (p *PlatformWatcherTron) WatchBlocks() {
	for {
		select {
		case <-p.done:
			return
		default:
			go p.FetchBlocks()
			p.Number++
			time.Sleep(3 * time.Second)
		}

	}
}

func (p *PlatformWatcherTron) WatchLatestBlock() {
	for {
		select {
		case <-p.done:
			return
		default:
			p.Service.GetNowBlock()
			time.Sleep(3 * time.Second)
		}
	}
}

func (p *PlatformWatcherTron) Stop() {
}

func NewPlatformWatcherTron(cfg *config.Config) PlatformWatcher {
	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())
	client.SetTimeout(60 * time.Second)

	log.Printf("Fetching USDT smart contract ...")
	usdtSmartContract, err := chain.GetUsdtSmartContract(client)
	log.Printf("USDT smart contract fetched")

	if err != nil {
		log.Fatalf("Error getting USDT smart contract: %v", err)
	}
	service := chain.NewTronBlockChainService(nil, usdtSmartContract, client)

	return &PlatformWatcherTron{
		Config:             cfg,
		done:               make(chan bool),
		Service:            service,
		ApplicationService: application.NewTransactionService(config.NewConfigWatchlistRepository(cfg.Wallets)),
		Number:             -1,
	}
}
