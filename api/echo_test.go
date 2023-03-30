package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Echo(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	admin := random.User()
	assert.Nil(t, c.CreateUser(admin))
	// Tokens
	token := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(admin.Claims()))))
	t.Run("ValidJWT", func(tt *testing.T) {
		expect.PUT(UserRoute).
			WithHeader("Authorization", token).
			WithJSON(random.User()).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("InvalidJWT", func(tt *testing.T) {
		expect.PUT(UserRoute).
			WithHeader("Authorization", "Bearer INVALID").
			WithJSON(random.User()).
			Expect().
			Status(http.StatusUnauthorized)
	})
}
