package web

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
)

type WSTradingListener struct {
	clients *socket.Clients
}

func (listener *WSTradingListener) BlockGenerated(block *model.Block) {
	//listener.clients.SendAll(block)
}

func NewWSTradingListener(clients *socket.Clients) *WSTradingListener {
	return &WSTradingListener{clients}
}
