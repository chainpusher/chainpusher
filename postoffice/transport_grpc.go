package postoffice

import "github.com/chainpusher/chainpusher/model"

type TransportGrpc struct {
}

func (po *TransportGrpc) Deliver(transaction *model.Transaction) error {
	return nil
}

func NewTransportGrpc() Transport {
	return &TransportGrpc{}
}