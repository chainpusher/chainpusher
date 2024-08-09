package gorm

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/counter"
	"gorm.io/gorm"
)

type CounterRepository struct {
	db *gorm.DB
}

func (repo *CounterRepository) FindCounterByKey(key string) (*counter.Counter, error) {
	var c counter.Counter
	r := repo.db.Where("key = ?", key).First(&c)
	if r.RowsAffected == 0 {
		return nil, nil
	}
	return &c, r.Error
}

func (repo *CounterRepository) IncrementCounterByKey(key string) (int64, error) {
	r := repo.
		db.
		Model(&counter.Counter{}).Where("key = ?", key).
		Update("value", gorm.Expr("value + ?", 1))
	if r.RowsAffected == 0 {
		return 0, nil
	}
	return r.RowsAffected, r.Error
}

func (repo *CounterRepository) Save(counter *counter.Counter) error {
	return repo.db.Save(counter).Error
}

func NewCounterRepository(db *gorm.DB) *CounterRepository {
	return &CounterRepository{db: db}
}
