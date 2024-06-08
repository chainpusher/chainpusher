package monitor

import (
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/chain"
	"github.com/sirupsen/logrus"
)

type PlatformWatcherEthereum struct {
	Done               chan bool
	Service            *chain.EthereumBlockChainService
	ApplicationService *application.TransactionService
	Number             *big.Int
	TimePollingCycle   time.Duration
	waitGroup          *sync.WaitGroup
}

func (p *PlatformWatcherEthereum) Start() {
	logrus.Println("Starting ethereum platform watcher ...")

	for {
		header, err := p.Service.GetNowBlock()
		if err != nil {
			log.Printf("Error getting block: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		logrus.Debugf("Latest block number is: %d", header.Number)
		p.Number = header.Number
		break
	}

	p.WatchBlocks()

}

func (p *PlatformWatcherEthereum) FetchBlocks(number *big.Int) {
	logrus.Debug("Fetching block: ", number)
	transactions, err := p.Service.GetTransactions(number)
	if err != nil {
		logrus.Warnf("Error getting block: %v", err)
		return
	}

	p.Number.Add(p.Number, big.NewInt(1))
	logrus.Infof("Block %d fetched with %d transactions", number, len(transactions))
	p.ApplicationService.AnalyzeTrade(transactions)
}

func (p *PlatformWatcherEthereum) WatchBlocks() {
	for {
		select {
		case <-p.Done:
			return
		default:
			go p.FetchBlocks(p.Number)

			time.Sleep(p.TimePollingCycle)
		}
	}
}

func (p *PlatformWatcherEthereum) WatchLatestBlock() {
	for {
		select {
		case <-p.Done:
			return
		default:
			p.Service.GetNowBlock()
			time.Sleep(15 * time.Second)
		}
	}
}

func (p *PlatformWatcherEthereum) Stop() {
	p.Done <- true
	close(p.Done)
}

func NewPlatformWatcherEthereum(
	timePollingCycle time.Duration,
	waitGroup *sync.WaitGroup,
	service *chain.EthereumBlockChainService,
	application *application.TransactionService) PlatformWatcher {

	return &PlatformWatcherEthereum{
		Done:               make(chan bool),
		Service:            service,
		ApplicationService: application,
		Number:             big.NewInt(-1),
		TimePollingCycle:   timePollingCycle,
		waitGroup:          waitGroup,
	}
}
