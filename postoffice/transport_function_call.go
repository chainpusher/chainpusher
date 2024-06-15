package postoffice

import "github.com/chainpusher/blockchain/model"

type TransportFunctionCall struct {
}

func (t *TransportFunctionCall) Deliver(_ *model.Block) error {
	return nil
}

func NewTransportFunctionCall() Transport {
	return &TransportFunctionCall{}
}
