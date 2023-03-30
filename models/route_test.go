package models

import (
	"testing"

	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&Vehicle{}, &Route{}))
	v := &Vehicle{
		Name:  "Ford Mustang",
		Plate: "COL000",
	}
	assert.Nil(t, db.Create(v).Error)
	t.Run("ValidRoute", func(tt *testing.T) {
		assert.Nil(tt, db.Create(&Route{
			VehicleUUID: v.UUID,
			Status:      RouteInProgress,
			StartNode:   1,
			EndNode:     5,
		}).Error)
	})
}
