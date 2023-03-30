/*
sqlite will be only used for unit tests and when no postgres DSN is provided
*/
package sqlite

import (
	"fmt"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var currentTestDB = 0

func NewTest() *gorm.DB {
	currentTestDB++
	return New(fmt.Sprintf("file:test-%d?mode=memory&cache=shared", currentTestDB))
}

func New(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	return db
}
