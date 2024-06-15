package postoffice

import "github.com/chainpusher/blockchain/model"

type TransportEmail struct {
}

func (po *TransportEmail) Deliver(_ *model.Block) error {
	return nil
}

func NewTransportEmail() Transport {
	return &TransportEmail{}
}
