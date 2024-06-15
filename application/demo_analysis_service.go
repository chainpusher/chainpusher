package application

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/config"
	model2 "github.com/chainpusher/chainpusher/model"

	"github.com/chainpusher/chainpusher/postoffice"
	"github.com/sirupsen/logrus"
)

type DemoAnalysisService struct {
	WatchlistRepository model2.WatchlistRepository

	postOffice postoffice.PostOffice

	WasDelivered bool
}

func NewTransactionDemoService(cfg *config.Config) *DemoAnalysisService {
	watchlist := config.NewConfigWatchlistRepository(cfg.Wallets)
	t := postoffice.NewTransportFromConfig(cfg)
	ps := postoffice.NewPostOfficeCoroutine(t)

	return &DemoAnalysisService{
		WatchlistRepository: watchlist,
		postOffice:          ps,
		WasDelivered:        false,
	}
}

func (t *DemoAnalysisService) AnalyzeTrade(block *model.Block) error {
	transactions := block.Transactions
	if !t.WasDelivered && len(transactions) > 0 {
		logrus.Info("DemoAnalysisService.AnalyzeTrade: Delivering transactions")
		aBlock := block.CloneWithTransactions(transactions)
		err := t.postOffice.Deliver(aBlock)
		if err != nil {
			return err
		}
		t.WasDelivered = true
	}

	return nil
}
