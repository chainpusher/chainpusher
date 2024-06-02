package postoffice

import "github.com/chainpusher/chainpusher/model"

type TransportHttp struct {
}

func (po *TransportHttp) Deliver(transaction *model.Transaction) error {
	return nil
}

func NewTransportHttp() Transport {
	return &TransportHttp{}
}
