package postoffice

import "github.com/chainpusher/blockchain/model"

type TransportGrpc struct {
}

func (po *TransportGrpc) Deliver(_ *model.Block) error {
	return nil
}

//func NewTransportGrpc() Transport {
//	return &TransportGrpc{}
//}
