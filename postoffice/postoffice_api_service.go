package postoffice

import "github.com/chainpusher/chainpusher/model"

type PostOfficeAPIService struct {
}

func (po *PostOfficeAPIService) Deliver(transactions []*model.Transaction) error {
	panic("API service not implemented yet")
}

func NewPostOfficeAPIService() PostOffice {
	return &PostOfficeAPIService{}
}
