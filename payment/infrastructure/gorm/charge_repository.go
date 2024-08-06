package gorm

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"gorm.io/gorm"
)

type ChargeRepository struct {
	db *gorm.DB
}

func (p *ChargeRepository) Save(entity *charge.Charge) error {
	if entity.ID > 0 {
		p.db.Updates(entity)
		return nil
	}
	r := p.db.Create(entity)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (p *ChargeRepository) Find(id int64) (*charge.Charge, error) {
	var c charge.Charge
	r := p.
		db.
		Preload("Pool").
		First(&c, id)
	if r.Error != nil {
		return nil, r.Error
	}

	return &c, nil
}

func NewChargeRepository(db *gorm.DB) *ChargeRepository {
	return &ChargeRepository{db: db}
}
