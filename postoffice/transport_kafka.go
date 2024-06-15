package postoffice

import "github.com/chainpusher/blockchain/model"

type TransportKafka struct {
}

func (t *TransportKafka) Deliver(_ *model.Block) error {
	return nil
}

func NewTransportKafka() Transport {
	return &TransportKafka{}
}
