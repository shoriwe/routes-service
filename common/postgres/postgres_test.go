package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgreSQL(t *testing.T) {
	dsn := "host=127.0.0.1 user=sulcud password=sulcud dbname=sulcud port=5432 sslmode=disable"
	db := New(dsn)
	conn, err := db.DB()
	assert.Nil(t, err)
	assert.Nil(t, conn.Close())
}
