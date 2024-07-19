package application

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
)

type TinyBlockServiceImpl struct {
	clients socket.Clients
}

func (svc *TinyBlockServiceImpl) Broadcast(block *model.Block) {
	event := &dto.JsonRpcEvent{
		Name: "block",
		Data: block,
	}
	svc.clients.Room("subscribe").Emit(event)
}

func (svc *TinyBlockServiceImpl) Subscribe(clientId int64) (socket.Client, error) {
	client, err := svc.clients.Get(clientId)
	if err != nil {
		return nil, err
	}
	err = svc.clients.Join(clientId, "subscribe")
	if err != nil {
		return client, err
	}

	return client, nil
}

func NewTinyBlockService(clients socket.Clients) TinyBlockService {
	return &TinyBlockServiceImpl{clients: clients}
}
