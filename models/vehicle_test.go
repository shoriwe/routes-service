package models

import (
	"testing"

	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestVehicle(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&Vehicle{}))
	t.Run("ValidVehicle", func(tt *testing.T) {
		assert.Nil(tt, db.Create(&Vehicle{
			Name:  "Ford Mustang",
			Plate: "COL000",
		}).Error)
		var v Vehicle
		assert.Nil(tt, db.Where("plate = ?", "COL000").First(&v).Error)
		assert.Equal(tt, "Ford Mustang", v.Name)
	})
}

func TestVehicle_BeforeSave(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&Vehicle{}))
	t.Run("NoName", func(tt *testing.T) {
		assert.NotNil(tt, db.Create(&Vehicle{
			Plate: "COL000",
		}).Error)
	})
	t.Run("InvalidPlate", func(tt *testing.T) {
		assert.NotNil(tt, db.Create(&Vehicle{
			Name:  "Ford Mustang",
			Plate: "######",
		}).Error)
	})
}
