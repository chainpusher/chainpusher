package facade

import (
	dto2 "github.com/chainpusher/chainpusher/trade/interfaces/price/facade/dto"
)

type ServiceFacade interface {
	GetPrice(crypto string) ([]*dto2.PriceDTO, error)

	LoadPrices() error
}
