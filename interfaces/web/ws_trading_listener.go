package web

import "github.com/chainpusher/blockchain/model"

type WSTradingListener struct {
	clients *Clients
}

func (listener *WSTradingListener) BlockGenerated(block *model.Block) {
	listener.clients.SendAll(block)
}

func NewWSTradingListener(clients *Clients) *WSTradingListener {
	return &WSTradingListener{clients}
}
