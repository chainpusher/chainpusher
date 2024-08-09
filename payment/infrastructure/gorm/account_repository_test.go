package gorm_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/test"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountRepository_FindBySecretKey(t *testing.T) {
	db := test.SetupTestDB()
	var a account.Account
	db.Preload("Secrets").First(&a)

	k := a.Secrets[0].Key
	assert.NotNil(t, k)

	a2, err := gorm.NewAccountRepository(db).FindBySecretKey(k)
	assert.Nil(t, err)
	assert.NotNil(t, a2)
}
