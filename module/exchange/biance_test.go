package exchange_test

import (
	"github.com/chainpusher/chainpusher/module/data"
	"github.com/chainpusher/chainpusher/module/exchange"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBianceExchangeService_GetPrice(t *testing.T) {
	svc := exchange.NewBianceExchangeService()
	prices, err := svc.GetPrice(data.Bitcoin)
	assert.Nil(t, err)
	assert.Len(t, prices, 1)
}
