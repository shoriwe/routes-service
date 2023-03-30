package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlate(t *testing.T) {
	for i := 0; i < 100; i++ {
		assert.Equal(t, 6, len(Plate()))
	}
}

func TestVehicle(t *testing.T) {
	assert.NotNil(t, Vehicle())
}
