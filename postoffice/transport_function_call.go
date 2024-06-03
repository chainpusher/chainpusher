package postoffice

import "github.com/chainpusher/chainpusher/model"

type TransportFunctionCall struct {
}

func (t *TransportFunctionCall) Deliver(transactions []*model.Transaction) error {
	return nil
}

func NewTransportFunctionCall() Transport {
	return &TransportFunctionCall{}
}
