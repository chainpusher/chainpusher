package gorm_test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/counter"
	"github.com/chainpusher/chainpusher/payment/domain/model/test"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounterRepository_FindCounterByKey(t *testing.T) {
	db := test.SetupTestDB()
	repo := gorm.NewCounterRepository(db)

	c := &counter.Counter{Key: "100", Value: 0}
	db.Create(c)

	affected, err := repo.IncrementCounterByKey("100")

	c, err = repo.FindCounterByKey("100")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), affected)
	assert.Equal(t, 1, c.Value)
	assert.Equal(t, int64(1), c.Id)
}
