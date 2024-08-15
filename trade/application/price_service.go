package application

import "github.com/chainpusher/chainpusher/trade/domain/model/price"

type PriceService interface {
	AddPrices(prices []*price.Price) ([]*price.Price, error)
}
