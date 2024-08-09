package gorm_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/test"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoolRepository_FindByAccountId(t *testing.T) {
	db := test.SetupTestDB()

	repo := gorm.NewPoolRepository(db)
	p, err := repo.FindByAccountId(1)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}
