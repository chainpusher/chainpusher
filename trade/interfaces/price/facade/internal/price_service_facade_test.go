package internal_test

import (
	"github.com/chainpusher/chainpusher/trade/application/impl"
	"github.com/chainpusher/chainpusher/trade/infrastructure/persistence/gorm"
	"github.com/chainpusher/chainpusher/trade/interfaces/price/facade/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriceServiceFacade_LoadPrices(t *testing.T) {
	db := gorm.SetupTestDatabase()
	repo := gorm.NewPriceRepository(db)
	priceService := impl.NewPriceService(repo)
	priceServiceFacade := internal.NewPriceServiceFacade(priceService)

	err := priceServiceFacade.LoadPrices()
	assert.Nil(t, err)
}
