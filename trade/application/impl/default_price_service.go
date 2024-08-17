package impl

import (
	"github.com/chainpusher/chainpusher/trade/application"
	"github.com/chainpusher/chainpusher/trade/domain/model/price"
)

type DefaultPriceService struct {
	repo price.Repository
}

func (d *DefaultPriceService) AddPrices(prices []*price.Price) ([]*price.Price, error) {
	return nil, nil
}

func NewDefaultPriceService(repo price.Repository) application.PriceService {
	return &DefaultPriceService{repo: repo}
}
