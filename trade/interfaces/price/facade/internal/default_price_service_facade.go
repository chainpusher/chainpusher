package internal

import (
	"github.com/chainpusher/chainpusher/module/exchange"
	"github.com/chainpusher/chainpusher/trade/application"
	dto2 "github.com/chainpusher/chainpusher/trade/interfaces/price/facade/dto"
)

type PriceServiceFacade struct {
	service application.PriceService

	exchange exchange.Service
}

func (f *PriceServiceFacade) GetPrice(crypto string) ([]*dto2.PriceDTO, error) {
	return nil, nil
}

func (f *PriceServiceFacade) LoadPrices() error {
	return nil
}

func NewPriceServiceFacade(service application.PriceService) *PriceServiceFacade {
	return &PriceServiceFacade{service: service}
}
