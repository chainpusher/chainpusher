package charge_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/test"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"testing"
)

func TestChargeRepository(t *testing.T) {
	db := test.SetupTestDB()
	var a = account.Account{}
	db.First(&a)

	assert.NotNil(t, a)

	repository := gorm.NewChargeRepository(db)
	c := &charge.Charge{
		Meta:      datatypes.JSON(`{"order_id": 1}`),
		AccountId: a.ID,
	}
	err := repository.Save(c)

	db.First(c)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), c.ID)
}
