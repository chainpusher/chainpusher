package impl_test

import (
	"testing"

	"github.com/chainpusher/chainpusher/trade/application/impl"
	"github.com/chainpusher/chainpusher/trade/infrastructure/persistence/gorm"
	impl2 "github.com/chainpusher/chainpusher/trade/interfaces/price/facade/impl"
	"github.com/stretchr/testify/assert"
)

func TestPriceServiceFacade_LoadPrices(t *testing.T) {
	db := gorm.SetupTestDatabase()
	repo := gorm.NewPriceRepository(db)
	priceService := impl.NewPriceService(repo)
	// priceServiceFacade := impl.NewPriceServiceFacade(priceService)
	priceServiceFacade := impl2.NewPriceServiceFacade(priceService)

	err := priceServiceFacade.LoadPrices()
	assert.Nil(t, err)
}
