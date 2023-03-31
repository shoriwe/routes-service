package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/controller"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateAPIKey(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	admin := random.User()
	admin.Type = models.Admin
	assert.Nil(t, c.CreateUser(admin))
	manager := random.User()
	manager.Type = models.Manager
	assert.Nil(t, c.CreateUser(manager))
	// Tokens
	adminToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(admin.Claims()))))
	managerToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(manager.Claims()))))
	t.Run("AsAdmin", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
		expect.
			PUT(APIKeyRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(&models.APIKey{
				VehicleUUID: vehicle.UUID,
			}).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
		expect.
			PUT(APIKeyRoute).
			WithHeader("Authorization", managerToken).
			WithJSON(&models.APIKey{
				VehicleUUID: vehicle.UUID,
			}).
			Expect().
			Status(http.StatusForbidden)
	})
	t.Run("NoJSON", func(tt *testing.T) {
		expect.PUT(APIKeyRoute).
			WithHeader("Authorization", adminToken).
			WithHeader("Content-Type", "application/json").
			WithText("{").
			Expect().
			Status(http.StatusBadRequest)
	})
}

func TestHandler_DeleteAPIKey(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	admin := random.User()
	admin.Type = models.Admin
	assert.Nil(t, c.CreateUser(admin))
	manager := random.User()
	manager.Type = models.Manager
	assert.Nil(t, c.CreateUser(manager))
	// Tokens
	adminToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(admin.Claims()))))
	managerToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(manager.Claims()))))
	t.Run("AsAdmin", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
		ak := &models.APIKey{
			VehicleUUID: vehicle.UUID,
		}
		assert.Nil(tt, c.CreateAPIKey(ak))
		expect.DELETE(APIKeyRoute+fmt.Sprintf("/%s", ak.UUID)).
			WithHeader("Authorization", adminToken).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		vehicle := random.Vehicle()
		assert.Nil(tt, c.CreateVehicle(vehicle))
		ak := &models.APIKey{
			VehicleUUID: vehicle.UUID,
		}
		assert.Nil(tt, c.CreateAPIKey(ak))
		expect.DELETE(APIKeyRoute+fmt.Sprintf("/%s", ak.UUID)).
			WithHeader("Authorization", managerToken).
			Expect().
			Status(http.StatusForbidden)
	})
}

func TestHandler_QueryAPIKeys(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	admin := random.User()
	admin.Type = models.Admin
	assert.Nil(t, c.CreateUser(admin))
	manager := random.User()
	manager.Type = models.Manager
	assert.Nil(t, c.CreateUser(manager))
	// Vehicles
	vehicle := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle))
	vehicle2 := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle2))
	// API keys
	for i := 0; i < 10; i++ {
		assert.Nil(t, c.CreateAPIKey(&models.APIKey{
			VehicleUUID: vehicle.UUID,
		}))
	}
	for i := 0; i < 10; i++ {
		assert.Nil(t, c.CreateAPIKey(&models.APIKey{
			VehicleUUID: vehicle2.UUID,
		}))
	}
	// Tokens
	adminToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(admin.Claims()))))
	managerToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(manager.Claims()))))
	t.Run("AsAdmin", func(tt *testing.T) {
		filter := &controller.VehicleFilter{Page: 1}
		expect.POST(APIKeyRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(filter).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		filter := &controller.VehicleFilter{Page: 1}
		expect.POST(APIKeyRoute).
			WithHeader("Authorization", managerToken).
			WithJSON(filter).
			Expect().
			Status(http.StatusForbidden)
	})
	t.Run("InvalidJSON", func(tt *testing.T) {
		expect.POST(APIKeyRoute).
			WithHeader("Authorization", adminToken).
			WithText("{").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
}
