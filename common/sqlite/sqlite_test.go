package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQLite(t *testing.T) {
	db := NewTest()
	conn, err := db.DB()
	assert.Nil(t, err)
	assert.Nil(t, conn.Close())
}
