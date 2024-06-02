package postoffice

import "github.com/chainpusher/chainpusher/model"

type TransportKafka struct {
}

func (t *TransportKafka) Deliver(transaction *model.Transaction) error {
	return nil
}

func NewTransportKafka() Transport {
	return &TransportKafka{}
}
