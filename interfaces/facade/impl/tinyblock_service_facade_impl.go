package impl

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/interfaces/facade"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
	"github.com/sirupsen/logrus"
)

type TinyBlockServiceFacadeImpl struct {
	service application.TinyBlockService
}

func (t *TinyBlockServiceFacadeImpl) Broadcast(block *model.Block) {
	t.service.Broadcast(block)
}

func (t *TinyBlockServiceFacadeImpl) Subscribe(clientId int64) {
	client, err := t.service.Subscribe(clientId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"clientId": clientId}).Error(err)
		return
	}

	err = client.Emit(&dto.JsonRpcEvent{
		Name: "subscribe",
		Data: nil,
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{"clientId": clientId}).Error(err)
	}
}

func (t *TinyBlockServiceFacadeImpl) GetTransactions(_ *dto.QueryTransactionsCommand) ([]*model.Transaction, error) {
	return nil, nil
}

func NewTinyBlockServiceFacade(service application.TinyBlockService) facade.TinyBlockServiceFacade {
	return &TinyBlockServiceFacadeImpl{service}
}
