package controller

import (
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestController_CreateAPIKey(t *testing.T) {
	c := NewTest()
	defer c.Close()
	t.Run("ValidVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
		assert.Nil(tt, c.CreateAPIKey(&models.APIKey{
			VehicleUUID: vehicle.UUID,
		}))
	})
	t.Run("RepeatedVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
		assert.Nil(tt, c.CreateAPIKey(&models.APIKey{
			VehicleUUID: vehicle.UUID,
		}))
		assert.Nil(tt, c.CreateAPIKey(&models.APIKey{
			VehicleUUID: vehicle.UUID,
		}))
	})
}

func TestController_DeleteAPIKey(t *testing.T) {
	c := NewTest()
	defer c.Close()
	vehicle := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle))
	t.Run("ValidVehicle", func(tt *testing.T) {
		ak := &models.APIKey{
			VehicleUUID: vehicle.UUID,
		}
		assert.Nil(tt, c.CreateAPIKey(ak))
		assert.Nil(tt, c.DeleteAPIKey(ak.UUID.String()))
	})
	t.Run("FakeVehicle", func(tt *testing.T) {
		ak := &models.APIKey{
			VehicleUUID: vehicle.UUID,
		}
		assert.Nil(tt, c.CreateAPIKey(ak))
		assert.Nil(tt, c.DeleteAPIKey("INVALID"))
	})
}

func TestController_QueryAPIKey(t *testing.T) {
	c := NewTest()
	defer c.Close()
	// Creating vehicles
	vehicle := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle))
	for i := 0; i < 10; i++ {
		assert.Nil(t, c.CreateAPIKey(
			&models.APIKey{
				VehicleUUID: vehicle.UUID,
			},
		))
	}
	vehicle2 := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle2))
	for i := 0; i < 10; i++ {
		assert.Nil(t, c.CreateAPIKey(
			&models.APIKey{
				VehicleUUID: vehicle2.UUID,
			},
		))
	}
	// Tests
	t.Run("QueryAll", func(tt *testing.T) {
		result, qErr := c.QueryAPIKeys(&APIKeyFilter{Page: 1})
		assert.Nil(tt, qErr)
		assert.NotNil(tt, result.Results)
	})
	t.Run("QueryVehicle-1", func(tt *testing.T) {
		vehicleUUID := vehicle.UUID.String()
		result, qErr := c.QueryAPIKeys(&APIKeyFilter{Page: 1, PageSize: 10, VehicleUUID: &vehicleUUID})
		assert.Nil(tt, qErr)
		assert.Equal(tt, 10, len(result.Results))
		assert.Equal(tt, int64(1), result.TotalPages)
		assert.NotNil(tt, result.Results)
	})
	t.Run("QueryVehicle-2", func(tt *testing.T) {
		vehicleUUID := vehicle2.UUID.String()
		result, qErr := c.QueryAPIKeys(&APIKeyFilter{Page: 1, PageSize: 10, VehicleUUID: &vehicleUUID})
		assert.Nil(tt, qErr)
		assert.Equal(tt, 10, len(result.Results))
		assert.Equal(tt, int64(1), result.TotalPages)
		assert.NotNil(tt, result.Results)
	})
}
