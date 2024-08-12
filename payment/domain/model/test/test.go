package test

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/counter"
	"github.com/chainpusher/chainpusher/payment/domain/model/price"
	"github.com/chainpusher/chainpusher/payment/domain/model/secret"
	"github.com/chainpusher/chainpusher/payment/domain/model/wallet"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupWalletPoolFixtures(db *gorm.DB, a *account.Account) error {

	a1 := wallet.Wallet{AccountId: a.ID, Blockchain: "ETHEREUM", Crypto: "USDT", Address: "1"}
	a2 := wallet.Wallet{AccountId: a.ID, Blockchain: "TRON", Crypto: "USDT", Address: "2"}

	db.Create(&a1)
	db.Create(&a2)
	return nil
}

func SetupFixtures(db *gorm.DB) error {
	app, err := account.NewAccount()
	if err != nil {
		return err
	}
	err = SetupWalletPoolFixtures(db, app)
	if err != nil {
		return err
	}
	db.Create(&app)

	return nil
}

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&charge.Charge{},
		&charge.Wallet{},
		&account.Account{},
		&secret.Secret{},
		&wallet.Wallet{},
		&price.Price{},
		&counter.Counter{},
	)
	if err != nil {
		panic("failed to migrate database")
	}

	err = SetupFixtures(db)
	if err != nil {
		panic("failed to setup fixtures")
	}

	return db.Debug()
}
