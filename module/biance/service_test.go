package biance_test

import (
	"github.com/chainpusher/chainpusher/module/biance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetPrice(t *testing.T) {
	svc := biance.NewService()
	prices, err := svc.GetPrice("BTCUSDT", "ETHUSDT")
	assert.Nil(t, err)
	assert.NotNil(t, prices)
	assert.Equal(t, 2, len(prices))
	assert.Equal(t, "BTCUSDT", prices[0].Symbol)
	assert.NotEmpty(t, prices[0].Price)
}
