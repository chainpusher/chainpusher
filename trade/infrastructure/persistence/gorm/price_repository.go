package gorm

import (
	"gorm.io/gorm"
)

type PriceRepository struct {
	db *gorm.DB
}

func (repo *PriceRepository) Save(price interface{}) error {
	return repo.db.Create(price).Error
}

func NewPriceRepository(db *gorm.DB) *PriceRepository {
	return &PriceRepository{db}
}
