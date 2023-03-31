package models

import (
	"testing"

	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestAPIKey(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&Vehicle{}, &APIKey{}))
	v := &Vehicle{
		Name:  "Ford Mustang",
		Plate: "COL000",
	}
	assert.Nil(t, db.Create(v).Error)
	t.Run("ValidAPIKey", func(tt *testing.T) {
		assert.Nil(tt, db.Create(&APIKey{
			VehicleUUID: v.UUID,
			Key:         "Test",
		}).Error)
	})
	t.Run("NoDefaultAPIKey", func(tt *testing.T) {
		assert.Nil(tt, db.Create(&APIKey{
			VehicleUUID: v.UUID,
		}).Error)
	})
}
