package postoffice

import "github.com/chainpusher/chainpusher/model"

type TransportTelgram struct {
}

func NewTransportTelgram() Transport {
	return &TransportTelgram{}
}

func (t *TransportTelgram) Deliver(transaction *model.Transaction) error {
	return nil
}
