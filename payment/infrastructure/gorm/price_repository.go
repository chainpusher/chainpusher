package gorm

import (
	"errors"
	"github.com/chainpusher/chainpusher/payment/domain/model/price"
	"gorm.io/gorm"
)

type PriceRepository struct {
	db *gorm.DB
}

func (repo *PriceRepository) FindPriceByAmount(amount int64) (*price.Price, error) {
	var p price.Price
	r := repo.db.Where("amount = ?", amount).First(&p)

	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if r.Error != nil {
		return nil, r.Error
	}
	return &p, nil
}

func NewPriceRepository(db *gorm.DB) *PriceRepository {
	return &PriceRepository{db: db}
}
