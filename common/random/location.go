package random

import (
	"math/rand"

	uuid "github.com/satori/go.uuid"
	"github.com/shoriwe/routes-service/models"
)

func Location(vehicleUUID uuid.UUID) *models.Location {
	return &models.Location{
		VehicleUUID: vehicleUUID,
		Latitude:    rand.Float64(),
		Longitude:   rand.Float64(),
		Speed:       rand.Float64(),
	}
}
