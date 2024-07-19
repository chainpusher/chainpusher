package impl

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/interfaces/facade"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
)

type TinyBlockServiceFacadeImpl struct {
	service application.TinyBlockService
}

func (t *TinyBlockServiceFacadeImpl) Broadcast(block *model.Block) {

}

func (t *TinyBlockServiceFacadeImpl) Subscribe(clientId int64) {
	client, err := t.service.Subscribe(clientId)
	if err != nil {
		return
	}

	client.Emit(&dto.JsonRpcResponseDto{
		Call: &dto.JsonRpcDto{
			Method: "subscribe",
			Params: nil,
		},
	})
}

func (t *TinyBlockServiceFacadeImpl) GetTransactions(command *dto.QueryTransactionsCommand) ([]*model.Transaction, error) {
	return nil, nil
}

func NewTinyBlockServiceFacade(service application.TinyBlockService) facade.TinyBlockServiceFacade {
	return &TinyBlockServiceFacadeImpl{service}
}
