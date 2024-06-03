package application

import (
	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
	"github.com/sirupsen/logrus"
)

type TransactionService struct {
	WatchlistRepository model.WatchlistRepository

	Postoffice postoffice.PostOffice
}

func NewTransactionService(repository model.WatchlistRepository) *TransactionService {
	return &TransactionService{
		WatchlistRepository: repository,
	}
}

func (t *TransactionService) AnalyzeTrade(transactions []*model.Transaction) error {
	logrus.Tracef("Transactions: %v", transactions)

	if watched := t.WatchlistRepository.In(transactions); len(watched) > 0 {
		t.Postoffice.Deliver(watched)
	}

	return nil
}
