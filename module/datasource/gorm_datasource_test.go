package datasource_test

import (
	"github.com/chainpusher/chainpusher/module/datasource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGormDatabase(t *testing.T) {
	db, err := datasource.GetGormDatabase(&datasource.Datasource{Driver: "sqlite3", Dsn: ":memory:"})
	assert.Nil(t, err)
	assert.NotNil(t, db)

	var one int
	db.Raw("SELECT 1").Scan(&one)
	assert.Equal(t, 1, one)
}
