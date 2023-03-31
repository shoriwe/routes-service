package random

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	assert.NotNil(t, Location(uuid.Nil))
}
