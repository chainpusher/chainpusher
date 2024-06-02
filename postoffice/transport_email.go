package postoffice

import "github.com/chainpusher/chainpusher/model"

type TransportEmail struct {
}

func (po *TransportEmail) Deliver(transaction *model.Transaction) error {
	return nil
}

func NewTransportEmail() Transport {
	return &TransportEmail{}
}
