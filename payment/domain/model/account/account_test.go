package account_test

import (
	app2 "github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(t *testing.T) {
	db := shared.SetupTestDB()

	account, err := app2.NewAccount()
	assert.Nil(t, err)
	assert.NotNil(t, account)
	db.Create(&account)

	assert.Less(t, int64(0), account.ID)
	assert.Less(t, int64(0), account.Secrets[0].Id)
}
