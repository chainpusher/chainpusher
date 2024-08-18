package assembler

import (
	"github.com/chainpusher/chainpusher/module/data"
	"github.com/chainpusher/chainpusher/module/exchange"
	"github.com/chainpusher/chainpusher/trade/domain/model/price"
)

type PriceDTOAssembler struct {
	cryptos *data.Pairs[string, string]
}

func (a *PriceDTOAssembler) ToPrice(dtoList []*exchange.Price) []*price.Price {
	prices := make([]*price.Price, 0)
	for _, d := range dtoList {
		prices = append(prices, a.toPrice(d))
	}
	return prices
}

func (a *PriceDTOAssembler) toPrice(p *exchange.Price) *price.Price {
	return &price.Price{
		Blockchain: a.cryptos.GetKey(p.Crypto),
		Crypto:     p.Crypto,
		Price:      p.Price,
	}
}

func NewPriceDTOAssembler() *PriceDTOAssembler {
	cryptos := data.NewPairs(
		data.NewPair(data.Bitcoin, data.Bitcoin),
		data.NewPair(data.ETH, data.Ethereum),
		data.NewPair(data.TRX, data.TRON),
	)
	return &PriceDTOAssembler{cryptos: &cryptos}
}
