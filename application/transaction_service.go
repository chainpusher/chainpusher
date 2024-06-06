package application

import (
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
	"github.com/sirupsen/logrus"
)

type TransactionService struct {
	WatchlistRepository model.WatchlistRepository

	Postoffice postoffice.PostOffice
}

func NewTransactionService(cfg *config.Config) *TransactionService {
	watchlist := config.NewConfigWatchlistRepository(cfg.Wallets)
	t := postoffice.NewTransportFromConfig(cfg)
	ps := postoffice.NewPostOfficeCoroutine(t)

	return &TransactionService{
		WatchlistRepository: watchlist,
		Postoffice:          ps,
	}
}

func (t *TransactionService) AnalyzeTrade(transactions []*model.Transaction) error {
	logrus.Tracef("Transactions: %v", transactions)

	if watched := t.WatchlistRepository.In(transactions); len(watched) > 0 {
		t.Postoffice.Deliver(watched)
	}

	return nil
}
