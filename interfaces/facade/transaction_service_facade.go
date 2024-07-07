package facade

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
)

type TransactionServiceFacade interface {
	GetTransactions(command *dto.QueryTransactionsCommand) ([]*model.Transaction, error)
}
