package datasource

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetGormDatabase(ds *Datasource) (*gorm.DB, error) {
	var dialector gorm.Dialector
	var db *gorm.DB
	var err error

	switch ds.Driver {
	case "sqlite3":
		dialector = sqlite.Open(ds.Dsn)
	case "mysql":
		dialector = mysql.Open(ds.Dsn)
	default:
		return nil, errors.New("unsupported driver")
	}

	db, err = gorm.Open(dialector, &gorm.Config{})
	return db, err
}
