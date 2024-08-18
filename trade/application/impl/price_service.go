package impl

import (
	"github.com/chainpusher/chainpusher/trade/application"
	"github.com/chainpusher/chainpusher/trade/domain/model/price"
)

type PriceService struct {
	repo price.Repository
}

func (d *PriceService) AddPrices(prices []*price.Price) ([]*price.Price, error) {
	if err := d.repo.Save(prices); err != nil {
		return nil, err
	}
	return prices, nil
}

func NewPriceService(repo price.Repository) application.PriceService {
	return &PriceService{repo: repo}
}
