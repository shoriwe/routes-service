package models

import (
	"testing"

	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&Vehicle{}, &Location{}))
	t.Run("ValidLocation", func(tt *testing.T) {
		v := &Vehicle{
			Name:  "Ford Mustang",
			Plate: "COL000",
		}
		assert.Nil(tt, db.Create(v).Error)
		assert.Nil(tt, db.Create(&Location{
			VehicleUUID: v.UUID,
			Latitude:    0,
			Longitude:   0,
			Speed:       50,
			Status:      EngineOn,
		}).Error)
		var l Location
		assert.Nil(tt, db.Where("vehicle_uuid = ?", v.UUID).First(&l).Error)
		assert.Equal(tt, float64(50), l.Speed)
	})
}
