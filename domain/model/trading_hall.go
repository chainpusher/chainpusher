package model

import "github.com/chainpusher/blockchain/model"

var tradingHall *TradingHall = NewTradingHall()

type TradingHall struct {
	listeners []TradingListener
}

func (h *TradingHall) AddListener(listener TradingListener) {
	h.listeners = append(h.listeners, listener)
}

func (h *TradingHall) NotifyListeners(block *model.Block) {
	for _, listener := range h.listeners {
		listener.BlockGenerated(block)
	}
}

func NewTradingHall() *TradingHall {
	return &TradingHall{}
}

func GetTradingHall() *TradingHall {
	return tradingHall
}
