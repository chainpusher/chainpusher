package web

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
)

type WSTradingListener struct {
	clients *socket.ClientsImpl
}

func (listener *WSTradingListener) BlockGenerated(block *model.Block) {
	//listener.clients.SendAll(block)
}

func NewWSTradingListener(clients *socket.ClientsImpl) *WSTradingListener {
	return &WSTradingListener{clients}
}
