package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	t.Run("ValidConnection", func(tt *testing.T) {
		dsn := "host=127.0.0.1 user=routes password=routes dbname=routes port=5432 sslmode=disable"
		db := New(dsn)
		conn, err := db.DB()
		assert.Nil(tt, err)
		assert.Nil(tt, conn.Close())
	})
	t.Run("InvalidConnection", func(tt *testing.T) {
		defer func() {
			err := recover()
			assert.NotNil(tt, err)
		}()
		New("")
	})
}
