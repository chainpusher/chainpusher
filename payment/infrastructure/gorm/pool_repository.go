package gorm

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/wallet"
	"gorm.io/gorm"
)

type PoolRepository struct {
	db *gorm.DB
}

func (r *PoolRepository) FindByAccountId(accountId int64) (*wallet.Pool, error) {
	var pool wallet.Pool
	r.
		db.
		Preload("Pool").
		Where("account_id = ?", accountId).First(&pool)
	return &pool, nil
}

func NewPoolRepository(db *gorm.DB) *PoolRepository {
	return &PoolRepository{db: db}
}
