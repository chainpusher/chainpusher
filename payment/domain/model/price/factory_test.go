package price_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/price"
	"github.com/chainpusher/chainpusher/payment/domain/model/test"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactory_AskForPrice(t *testing.T) {
	db := test.SetupTestDB()
	f := price.NewFactory(gorm.NewCounterRepository(db), gorm.NewPriceRepository(db))

	p, err := f.AskForPrice(100)
	assert.Nil(t, err)
	assert.Equal(t, int64(100), p.Amount)
	assert.Equal(t, 0, p.State)
	assert.Equal(t, true, p.Used)

	p, err = f.AskForPrice(100)
	assert.Nil(t, err)
	assert.Equal(t, int64(100), p.Amount)
	assert.Equal(t, 1, p.State)
	assert.Equal(t, true, p.Used)
}
