package exchange

import (
	"github.com/chainpusher/chainpusher/module/biance"
)

type BianceExchangeService struct {
	service *biance.Service

	assembler *BianceServiceAssembler
}

func (svc *BianceExchangeService) GetPrice(cryptos ...string) ([]*Price, error) {
	var symbols []string
	var prices1 []*biance.Price
	var prices2 []*Price
	var err error

	if symbols, err = svc.assembler.ToSymbols(cryptos); err != nil {
		return nil, err
	}

	if prices1, err = svc.service.GetPrice(symbols...); err != nil {
		return nil, err
	}

	if prices2, err = svc.assembler.ToPrices(prices1); err != nil {
		return nil, err
	}

	return prices2, nil
}

func NewBianceExchangeService() *BianceExchangeService {
	symbols := NewSymbols()
	assembler := NewBianceServiceAssembler(&symbols)
	return &BianceExchangeService{service: biance.NewService(), assembler: assembler}
}
