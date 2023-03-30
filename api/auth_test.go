package api

import (
	"net/http"
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Login(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	user := random.User()
	assert.Nil(t, c.CreateUser(user))
	t.Run("ValidCredentials", func(tt *testing.T) {
		expect.POST(LoginRoute).
			WithJSON(&models.User{
				Username: user.Username,
				Password: user.Password,
			}).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Value("jwt").String().Length().Gt(0)
	})
	t.Run("InvalidCredentials", func(tt *testing.T) {
		expect.POST(LoginRoute).
			WithJSON(&models.User{
				Username: user.Username,
				Password: "INVALID",
			}).
			Expect().
			Status(http.StatusUnauthorized)
	})
}
