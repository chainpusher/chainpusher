package application

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
)

type TinyBlockService interface {
	Subscribe(clientId int64) (socket.Client, error)

	Unsubscribe(clientId int64) (socket.Client, error)

	Broadcast(block *model.Block)
}
