package impl

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/facade"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
)

type TransactionServiceFacadeImpl struct {
}

func (t *TransactionServiceFacadeImpl) GetTransactions(command *dto.QueryTransactionsCommand) ([]*model.Transaction, error) {
	return nil, nil
}

func NewTransactionServiceFacadeImpl() facade.TransactionServiceFacade {
	return &TransactionServiceFacadeImpl{}
}
