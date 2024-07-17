package impl

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/facade"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
)

type TinyBlockServiceFacadeImpl struct {
}

func (t *TinyBlockServiceFacadeImpl) Subscribe(client int) {

}

func (t *TinyBlockServiceFacadeImpl) GetTransactions(command *dto.QueryTransactionsCommand) ([]*model.Transaction, error) {
	return nil, nil
}

func NewTransactionServiceFacadeImpl() facade.TinyBlockServiceFacade {
	return &TinyBlockServiceFacadeImpl{}
}
