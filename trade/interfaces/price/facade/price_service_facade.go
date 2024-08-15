package facade

import (
	dto2 "github.com/chainpusher/chainpusher/trade/interfaces/price/facade/dto"
)

type ServiceFacade interface {
	GetPrice(crypto string) ([]*dto2.PriceDTO, error)

	StorePrice(prices *[]dto2.PriceCommand) ([]*dto2.PriceCommand, error)
}
