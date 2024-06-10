package monitor

import (
	"log"
	"os"
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
	ApplicationService application.TransactionService
	Number             int64
	isOneTime          bool
}

func (p *PlatformWatcherTron) Start() {
	for {
		latest, transactions, err := p.Service.GetNowBlock()
		if err != nil {
			logrus.Errorf("Error getting block: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		logrus.Debugf("Latest block number: %d and transactions %d", latest.BlockHeader.RawData.Number, len(transactions))
		p.ApplicationService.AnalyzeTrade(transactions)
		p.Number = latest.BlockHeader.RawData.Number + 1
		break
	}

	time.Sleep(3 * time.Second)

	p.WatchBlocks()

	log.Println("Starting Tron platform watcher ...")
}

func (p *PlatformWatcherTron) FetchBlocks(number int64) {
	logrus.Debug("Fetching block ", number)
	transactions, err := p.Service.GetBlock(number)
	if err != nil {
		logrus.Warnf("Error getting block: %v", err)
		return
	}
	logrus.Debugf("Block %d fetched with %d transactions", number, len(transactions))
	p.ApplicationService.AnalyzeTrade(transactions)
	p.Number++
}

func (p *PlatformWatcherTron) WatchBlocks() {
	for {
		select {
		case <-p.done:
			return
		default:
			go p.FetchBlocks(p.Number)

			if p.isOneTime {
				time.Sleep(5 * time.Second)
				os.Exit(0)
			}

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
	p.done <- true
}

func NewPlatformWatcherTron(
	cfg *config.Config,
	application application.TransactionService,
	channel chan interface{},
	isOneTime bool,
) PlatformWatcher {

	client := client.NewGrpcClient("")
	client.Start(grpc.WithInsecure())
	client.SetTimeout(60 * time.Second)

	logrus.Info("Fetching USDT smart contract ...")
	usdtSmartContract, err := chain.GetUsdtSmartContract(client)
	logrus.Info("USDT smart contract fetched")

	if err != nil {
		log.Fatalf("Error getting USDT smart contract: %v", err)
	}
	service := chain.NewTronBlockChainService(nil, usdtSmartContract, client, channel)

	return &PlatformWatcherTron{
		Config:             cfg,
		done:               make(chan bool),
		Service:            service,
		ApplicationService: application,
		Number:             -1,
		isOneTime:          isOneTime,
	}
}
