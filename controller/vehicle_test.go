package controller

import (
	"fmt"
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestController_CreateVehicle(t *testing.T) {
	c := NewTest()
	defer c.Close()
	t.Run("ValidVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
	})
	t.Run("RepeatedVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
		assert.NotNil(tt, c.CreateVehicle(vehicle))
	})
}

func TestController_DeleteVehicle(t *testing.T) {
	c := NewTest()
	defer c.Close()
	t.Run("DeleteExistingVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(t, c.CreateVehicle(vehicle))
		assert.Nil(tt, c.DeleteVehicle(vehicle.UUID.String()))
	})
	t.Run("DeleteNonExistingVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(t, c.CreateVehicle(vehicle))
		assert.Nil(tt, c.DeleteVehicle("UUID"))
	})
}

func TestController_UpdateVehicle(t *testing.T) {
	c := NewTest()
	defer c.Close()
	t.Run("UpdateVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(t, c.CreateVehicle(vehicle))
		var vehicle2 models.Vehicle = *vehicle
		vehicle2.Name = "New Name"
		assert.Nil(tt, c.UpdateVehicle(&vehicle2))
	})
}

func TestController_QueryVehicle(t *testing.T) {
	c := NewTest()
	defer c.Close()
	for i := 0; i < 10; i++ {
		vehicle := random.Vehicle()
		vehicle.Name = fmt.Sprintf("sedan-%s", vehicle.Name)
		assert.Nil(t, c.CreateVehicle(vehicle))
	}
	for i := 0; i < 10; i++ {
		vehicle := random.Vehicle()
		vehicle.Name = fmt.Sprintf("sport-%s", vehicle.Name)
		assert.Nil(t, c.CreateVehicle(vehicle))
	}
	t.Run("QueryAll", func(tt *testing.T) {
		results, qErr := c.QueryVehicles(&VehicleFilter{Page: 1})
		assert.Nil(tt, qErr)
		assert.Equal(tt, DefaultPageSize, len(results.Results))
		assert.Equal(tt, 20/DefaultPageSize, int(results.TotalPages))
	})
	t.Run("QueryPlate", func(tt *testing.T) {
		plate := "A%"
		results, qErr := c.QueryVehicles(&VehicleFilter{Page: 1, Plate: &plate})
		assert.Nil(tt, qErr)
		assert.NotNil(tt, results.Results)
	})
	t.Run("QuerybyName", func(tt *testing.T) {
		name := "sedan-%"
		results, qErr := c.QueryVehicles(&VehicleFilter{Page: 1, Name: &name})
		assert.Nil(tt, qErr)
		assert.NotNil(tt, results.Results)
	})
}
