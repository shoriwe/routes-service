package controller

import (
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestController_Login(t *testing.T) {
	c := NewTest()
	defer c.Close()
	user := random.User()
	assert.Nil(t, c.CreateUser(user))
	t.Run("ValidCredentials", func(tt *testing.T) {
		token, lErr := c.Login(user)
		assert.Nil(tt, lErr)
		assert.Greater(tt, len(token), 0)
	})
	t.Run("InvalidCredentials", func(tt *testing.T) {
		token, lErr := c.Login(random.User())
		assert.NotNil(tt, lErr)
		assert.Equal(tt, 0, len(token))
	})
	t.Run("InvalidPassword", func(tt *testing.T) {
		var user2 models.User = *user
		user2.Password = "invalid"
		token, lErr := c.Login(&user2)
		assert.NotNil(tt, lErr)
		assert.Equal(tt, 0, len(token))
	})
}

func TestController_Authorize(t *testing.T) {
	c := NewTest()
	defer c.Close()
	user := random.User()
	assert.Nil(t, c.CreateUser(user))
	t.Run("ValidCredentials", func(tt *testing.T) {
		token, lErr := c.Login(user)
		assert.Nil(tt, lErr)
		assert.Greater(tt, len(token), 0)
		creds, aErr := c.AuthorizeUser(token)
		assert.Nil(tt, aErr)
		assert.Equal(tt, user.UUID, creds.UUID)
	})
	t.Run("InvalidCredentials", func(tt *testing.T) {
		_, aErr := c.AuthorizeUser(c.Session.New(random.User().Claims()))
		assert.NotNil(tt, aErr)
	})
	t.Run("InvalidJWT", func(tt *testing.T) {
		_, aErr := c.AuthorizeUser("INVALID")
		assert.NotNil(tt, aErr)
	})
}
