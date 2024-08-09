package gorm_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/price"
	"github.com/chainpusher/chainpusher/payment/domain/model/test"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriceRepository_FindPriceByAmount(t *testing.T) {
	db := test.SetupTestDB()

	repo := gorm.NewPriceRepository(db)

	p := &price.Price{Amount: 100}
	c := charge.Charge{Price: *p}
	db.Create(&c)

	p, err := repo.FindPriceByAmount(100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}
