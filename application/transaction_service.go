package application

import (
	"log"

	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/postoffice"
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

func (t *TransactionService) AnalyzeTrade(transaction *model.Transaction) error {
	log.Printf("Transfer: %v", transaction.Logging())

	if isOn := t.WatchlistRepository.IsOnList(transaction.Payee); !isOn {
		return nil
	}

	t.Postoffice.Deliver(transaction)

	return nil
}

func (t *TransactionService) AnalyzeTrades(transaction []*model.Transaction) {
	for _, trade := range transaction {
		t.AnalyzeTrade(trade)
	}
}
