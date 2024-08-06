package internal_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	dto2 "github.com/chainpusher/chainpusher/payment/interfaces/charge/facade/dto"
	"github.com/chainpusher/chainpusher/payment/interfaces/charge/facade/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharge(t *testing.T) {
	db := shared.SetupTestDB()
	f := internal.NewChargeServiceFacade(db)
	repo := gorm.NewChargeRepository(db)

	dto := &dto2.CreateChargeDTO{}
	c, err := f.Charge(dto)
	assert.Nil(t, err)
	assert.NotNil(t, c)

	f.Charged(c.Id)

	c2, err := repo.Find(c.Id)
	assert.Nil(t, err)
	assert.NotNil(t, c2)
	assert.Equal(t, charge.PAID, c2.Status)
}
