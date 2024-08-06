package charge_test

import (
	app2 "github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"testing"
)

func TestChargeRepository(t *testing.T) {
	db := shared.SetupTestDB()
	var app = app2.Account{}
	db.First(&app)
	db.Create(&app)

	repository := gorm.NewChargeRepository(db)
	c := &charge.Charge{
		Meta:      datatypes.JSON(`{"order_id": 1}`),
		AccountId: app.ID,
	}
	err := repository.Save(c)

	db.First(c)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), c.ID)
}
