package internal

import (
	"github.com/chainpusher/chainpusher/module/data"
	"github.com/chainpusher/chainpusher/module/exchange"
	"github.com/chainpusher/chainpusher/trade/application"
	"github.com/chainpusher/chainpusher/trade/domain/model/price"
	dto2 "github.com/chainpusher/chainpusher/trade/interfaces/price/facade/dto"
	"github.com/chainpusher/chainpusher/trade/interfaces/price/facade/internal/assembler"
	"github.com/sirupsen/logrus"
)

type PriceServiceFacade struct {
	service   application.PriceService
	exchange  exchange.Service
	assembler *assembler.PriceDTOAssembler
}

func (f *PriceServiceFacade) GetPrice(crypto string) ([]*dto2.PriceDTO, error) {
	return nil, nil
}

func (f *PriceServiceFacade) LoadPrices() error {
	var p []*exchange.Price
	var p3 []*price.Price
	var err error
	if p, err = f.exchange.GetPrice(data.Bitcoin, data.ETH, data.TRX); err != nil {
		return err
	}

	p2 := f.assembler.ToPrice(p)
	if p3, err = f.service.AddPrices(p2); err != nil {
		return err
	}
	logrus.Infof("Prices loaded: %v", p3)
	return nil
}

func NewPriceServiceFacade(service application.PriceService) *PriceServiceFacade {
	return &PriceServiceFacade{
		service:   service,
		exchange:  exchange.NewBianceExchangeService(),
		assembler: assembler.NewPriceDTOAssembler(),
	}
}
