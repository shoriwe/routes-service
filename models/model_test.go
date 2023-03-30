package models

import (
	"testing"

	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/stretchr/testify/assert"
)

/*
TestBaseModel basic unit test to reduce the coverage footprint
*/
func TestBaseModel(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&Model{}))
	assert.Nil(t, db.Create(&Model{}).Error)
}
