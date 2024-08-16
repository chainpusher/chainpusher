package exchange_test

import (
	"github.com/chainpusher/chainpusher/module/data"
	"github.com/chainpusher/chainpusher/module/exchange"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSymbols(t *testing.T) {
	symbols := exchange.NewSymbols()

	assert.Equal(t, "BTCUSDT", symbols.GetValue(data.Bitcoin))
	assert.Equal(t, "ETHUSDT", symbols.GetValue(data.ETH))
	assert.Equal(t, "TRXUSDT", symbols.GetValue(data.TRX))

	assert.Equal(t, symbols.GetKey("BTCUSDT"), data.Bitcoin)
	assert.Equal(t, symbols.GetKey("ETHUSDT"), data.ETH)
	assert.Equal(t, symbols.GetKey("TRXUSDT"), data.TRX)
}
