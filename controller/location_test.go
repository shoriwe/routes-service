package controller

import (
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestController_LocationListener(t *testing.T) {
	c := NewTest()
	defer c.Close()
	vehicle := random.Vehicle()
	assert.Nil(t, c.CreateVehicle(vehicle))
	w, pErr := c.LocationProducer(vehicle.UUID.String())
	assert.Nil(t, pErr)
	defer w.Close()
	ch, lErr := c.LocationListener(vehicle.UUID.String())
	assert.Nil(t, lErr)
	locations := make([]*models.Location, 0, 100)
	for i := 0; i < 100; i++ {
		location := random.Location(vehicle.UUID)
		assert.Nil(t, w.Send(location))
		locations = append(locations, location)
	}
	for _, l := range locations {
		assert.Equal(t, l, <-ch)
	}
}
