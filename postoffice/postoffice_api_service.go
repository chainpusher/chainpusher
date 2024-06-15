package postoffice

import "github.com/chainpusher/blockchain/model"

type APIService struct {
}

func (po *APIService) Deliver(block *model.Block) error {
	panic("API service not implemented yet")
}

func (po *APIService) GetTransports() []Transport {
	return []Transport{}
}

func NewAPIService() PostOffice {
	return &APIService{}
}
