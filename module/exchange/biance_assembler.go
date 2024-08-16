package exchange

import (
	"fmt"
	"github.com/chainpusher/chainpusher/module/biance"
	"github.com/chainpusher/chainpusher/module/data"
	"strconv"
	"strings"
)

type BianceServiceAssembler struct {
	symbols *data.Pairs[string, string]
}

func (a *BianceServiceAssembler) ToSymbols(cryptos []string) ([]string, error) {
	symbols := make([]string, 0, len(cryptos))
	for _, crypto := range cryptos {
		symbol := a.symbols.GetValue(crypto)
		if symbol == "" {
			return nil, fmt.Errorf("symbol not found for crypto %s", crypto)
		}
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

func (a *BianceServiceAssembler) ToPrices(prices []*biance.Price) ([]*Price, error) {
	var result []*Price
	for _, price := range prices {
		if p, err := a.ToPrice(price); err != nil {
			return nil, err
		} else {
			result = append(result, p)
		}
	}
	return result, nil
}

func (a *BianceServiceAssembler) ToPrice(price *biance.Price) (*Price, error) {
	p, err := strconv.ParseInt(strings.Replace(price.Price, ".", "", -1), 10, 64)
	if err != nil {
		return nil, err
	}

	crypto := a.symbols.GetKey(price.Symbol)

	return &Price{
		Crypto: crypto,
		Price:  p,
	}, nil
}

func NewBianceServiceAssembler(symbols *data.Pairs[string, string]) *BianceServiceAssembler {
	return &BianceServiceAssembler{symbols: symbols}
}
