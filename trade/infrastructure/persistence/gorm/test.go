package gorm

import (
	"github.com/chainpusher/chainpusher/trade/domain/model/price"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&price.Price{},
	)

	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
