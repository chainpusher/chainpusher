package postoffice

import "github.com/chainpusher/chainpusher/model"

type TransportRabbitMQ struct {
}

func (t *TransportRabbitMQ) Deliver(transactions []*model.Transaction) error {
	panic("RabbitMQ not implemented yet")
	return nil
}

func NewTransportRabbitMQ() Transport {
	return &TransportRabbitMQ{}
}
