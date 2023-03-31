package api

import (
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestLocationListener(t *testing.T) {
	c, baseURL, serverClose := NewTestWS(t)
	defer serverClose()
	// User
	user := random.User()
	assert.Nil(t, c.CreateUser(user))
	jwtToken, _ := c.Login(user)
	// Vehicle
	vehicle := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle))
	// API Key
	ak := &models.APIKey{VehicleUUID: vehicle.UUID}
	assert.Nil(t, c.CreateAPIKey(ak))
	// Producer
	w, wErr := c.LocationProducer(vehicle.UUID.String())
	assert.Nil(t, wErr)
	defer w.Close()
	locations := make([]*models.Location, 0, 5)
	for i := 0; i < 5; i++ {
		locations = append(locations, random.Location(vehicle.UUID))
	}
	// Consumer
	conn, _, dErr := websocket.DefaultDialer.Dial(baseURL+LocationConsumerRoute, http.Header{})
	assert.Nil(t, dErr)
	defer conn.Close()
	assert.Nil(t, conn.WriteJSON(LocationHandshake{Token: jwtToken}))
	var hs LocationHandshake
	assert.Nil(t, conn.ReadJSON(&hs))
	assert.True(t, hs.Status.Succeed)
	assert.Nil(t, conn.WriteJSON(LocationHandshake{VehicleUUID: vehicle.UUID.String()}))
	assert.Nil(t, conn.ReadJSON(&hs))
	// Produce
	go func() {
		for _, l := range locations {
			w.Send(l)
		}
	}()
	// Consume
	for _, l := range locations {
		rErr := conn.ReadJSON(&hs)
		assert.Nil(t, rErr)
		assert.Equal(t, l.Latitude, hs.Location.Latitude)
		assert.Equal(t, l.Longitude, hs.Location.Longitude)
		assert.Equal(t, l.Speed, hs.Location.Speed)
	}
}

func TestLocationProducer(t *testing.T) {
	c, baseURL, serverClose := NewTestWS(t)
	defer serverClose()
	// Vehicle
	vehicle := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle))
	// API Key
	ak := &models.APIKey{VehicleUUID: vehicle.UUID}
	assert.Nil(t, c.CreateAPIKey(ak))
	// Producer
	conn, _, dErr := websocket.DefaultDialer.Dial(baseURL+LocationProducerRoute, http.Header{})
	assert.Nil(t, dErr)
	defer conn.Close()
	assert.Nil(t, conn.WriteJSON(LocationHandshake{Token: ak.Key}))
	var hs LocationHandshake
	assert.Nil(t, conn.ReadJSON(&hs))
	assert.True(t, hs.Status.Succeed)
	//
	// Consumer
	ch, lErr := c.LocationListener(vehicle.UUID.String())
	assert.Nil(t, lErr)
	// Produce
	locations := make([]*models.Location, 0, 2)
	for i := 0; i < 2; i++ {
		location := random.Location(vehicle.UUID)
		locations = append(locations, location)
		assert.Nil(t, conn.WriteJSON(LocationHandshake{Location: location}))
	}
	// Consume
	for _, l := range locations {
		loc := <-ch
		assert.Equal(t, l.Latitude, loc.Latitude)
		assert.Equal(t, l.Longitude, loc.Longitude)
		assert.Equal(t, l.Speed, loc.Speed)
	}
}
