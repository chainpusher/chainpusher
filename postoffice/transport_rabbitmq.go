package postoffice

import "github.com/chainpusher/blockchain/model"

type TransportRabbitMQ struct {
}

func (t *TransportRabbitMQ) Deliver(_ *model.Block) error {
	panic("RabbitMQ not implemented yet")
	return nil
}

func NewTransportRabbitMQ() Transport {
	return &TransportRabbitMQ{}
}
