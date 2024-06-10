package application

import "github.com/chainpusher/chainpusher/model"

type TransactionService interface {
	AnalyzeTrade(transactions []*model.Transaction) error
}
