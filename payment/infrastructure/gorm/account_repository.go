package gorm

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func (repository *AccountRepository) FindBySecretKey(secretKey string) (*account.Account, error) {
	var a account.Account
	repository.
		db.
		Preload("Pool").
		Preload("Pool.Wallets").
		Preload("Secrets", "key = ?", secretKey).
		First(&a)
	return &a, nil
}

func NewAccountRepository(db *gorm.DB) account.Repository {
	return &AccountRepository{
		db: db,
	}

}
