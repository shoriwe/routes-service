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

func TestHandler_CreateVehicle(t *testing.T) {
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
		expect.
			PUT(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(random.Vehicle()).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		expect.
			PUT(VehicleRoute).
			WithHeader("Authorization", managerToken).
			WithJSON(random.Vehicle()).
			Expect().
			Status(http.StatusForbidden)
	})
	t.Run("NoJSON", func(tt *testing.T) {
		expect.PUT(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithHeader("Content-Type", "application/json").
			WithText("{").
			Expect().
			Status(http.StatusBadRequest)
	})
	t.Run("RepeatedVehicle", func(tt *testing.T) {
		vehicle := random.Vehicle()
		expect.PUT(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(vehicle).
			Expect().
			Status(http.StatusOK)
		expect.PUT(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(vehicle).
			Expect().
			Status(http.StatusInternalServerError)
	})
}

func TestHandler_DeleteVehicle(t *testing.T) {
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
		c.CreateVehicle(vehicle)
		expect.DELETE(VehicleRoute+fmt.Sprintf("/%s", vehicle.UUID)).
			WithHeader("Authorization", adminToken).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		vehicle := random.Vehicle()
		c.CreateVehicle(vehicle)
		expect.DELETE(VehicleRoute+fmt.Sprintf("/%s", vehicle.UUID)).
			WithHeader("Authorization", managerToken).
			Expect().
			Status(http.StatusForbidden)
	})
}

func TestHandler_UpdateVehicle(t *testing.T) {
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
		c.CreateVehicle(vehicle)
		vehicle.Plate = "NEW000"
		expect.PATCH(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(vehicle).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		vehicle := random.Vehicle()
		c.CreateVehicle(vehicle)
		vehicle.Plate = "NEW000"
		expect.PATCH(VehicleRoute).
			WithHeader("Authorization", managerToken).
			WithJSON(vehicle).
			Expect().
			Status(http.StatusForbidden)
	})
	t.Run("InvalidJSON", func(tt *testing.T) {
		vehicle := random.Vehicle()
		c.CreateVehicle(vehicle)
		vehicle.Plate = "NEW000"
		expect.PATCH(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithText("{").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
}

func TestHandler_QueryVehicles(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	admin := random.User()
	admin.Type = models.Admin
	assert.Nil(t, c.CreateUser(admin))
	manager := random.User()
	manager.Type = models.Manager
	assert.Nil(t, c.CreateUser(manager))
	// Vehicles
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
	// Tokens
	adminToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(admin.Claims()))))
	managerToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(manager.Claims()))))
	t.Run("AsAdmin", func(tt *testing.T) {
		filter := &controller.VehicleFilter{Page: 1}
		expect.POST(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(filter).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		filter := &controller.VehicleFilter{Page: 1}
		expect.POST(VehicleRoute).
			WithHeader("Authorization", managerToken).
			WithJSON(filter).
			Expect().
			Status(http.StatusForbidden)
	})
	t.Run("InvalidJSON", func(tt *testing.T) {
		expect.POST(VehicleRoute).
			WithHeader("Authorization", adminToken).
			WithText("{").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
}
