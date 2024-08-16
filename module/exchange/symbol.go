package exchange

import "github.com/chainpusher/chainpusher/module/data"

type Symbols data.Pairs[string, string]

func NewSymbols() data.Pairs[string, string] {
	return data.NewPairs[string, string](
		data.NewPair(data.Bitcoin, "BTCUSDT"),
		data.NewPair(data.ETH, "ETHUSDT"),
		data.NewPair(data.TRX, "TRXUSDT"),
	)
}
