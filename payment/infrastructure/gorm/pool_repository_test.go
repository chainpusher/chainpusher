package gorm_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoolRepository_FindByAccountId(t *testing.T) {
	db := shared.SetupTestDB()

	repo := gorm.NewPoolRepository(db)
	p, err := repo.FindByAccountId(1)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}
