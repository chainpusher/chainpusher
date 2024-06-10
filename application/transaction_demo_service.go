package application

import (
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
	"github.com/sirupsen/logrus"
)

type TransactionDemoService struct {
	WatchlistRepository model.WatchlistRepository

	Postoffice postoffice.PostOffice

	WasDelivered bool
}

func NewTransactionDemoService(cfg *config.Config) *TransactionDemoService {
	watchlist := config.NewConfigWatchlistRepository(cfg.Wallets)
	t := postoffice.NewTransportFromConfig(cfg)
	ps := postoffice.NewPostOfficeCoroutine(t)

	return &TransactionDemoService{
		WatchlistRepository: watchlist,
		Postoffice:          ps,
		WasDelivered:        false,
	}
}

func (t *TransactionDemoService) AnalyzeTrade(transactions []*model.Transaction) error {

	if !t.WasDelivered && len(transactions) > 0 {
		logrus.Info("TransactionDemoService.AnalyzeTrade: Delivering transactions")
		t.Postoffice.Deliver([]*model.Transaction{transactions[0]})
		t.WasDelivered = true
	}

	return nil
}
