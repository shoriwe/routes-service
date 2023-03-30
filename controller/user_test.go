package controller

import (
	"fmt"
	"testing"

	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/models"
	"github.com/stretchr/testify/assert"
)

func TestController_CreateUser(t *testing.T) {
	c := NewTest()
	defer c.Close()
	t.Run("ValidCredentials", func(tt *testing.T) {
		user := random.User()
		assert.Nil(tt, c.CreateUser(user))
	})
	t.Run("RepeatedUser", func(tt *testing.T) {
		user := random.User()
		assert.Nil(tt, c.CreateUser(user))
		assert.NotNil(tt, c.CreateUser(user))
	})
}

func TestController_DeleteUser(t *testing.T) {
	c := NewTest()
	defer c.Close()
	t.Run("DeleteExistingUser", func(tt *testing.T) {
		user := random.User()
		assert.Nil(t, c.CreateUser(user))
		assert.Nil(tt, c.DeleteUser(user.UUID.String()))
	})
	t.Run("DeleteNonExistingUser", func(tt *testing.T) {
		user := random.User()
		assert.Nil(t, c.CreateUser(user))
		assert.Nil(tt, c.DeleteUser("UUID"))
	})
}

func TestController_UpdateUser(t *testing.T) {
	c := NewTest()
	defer c.Close()
	t.Run("UpdateUser", func(tt *testing.T) {
		user := random.User()
		user.Type = models.Admin
		assert.Nil(t, c.CreateUser(user))
		var user2 models.User = *user
		user2.Type = models.Manager
		user2.Password = ""
		user2.PasswordHash = nil
		assert.Nil(tt, c.UpdateUser(&user2))
	})
}

func TestController_QueryUser(t *testing.T) {
	c := NewTest()
	defer c.Close()
	for i := 0; i < 10; i++ {
		user := random.User()
		user.Username = fmt.Sprintf("admin-%s", user.Username)
		user.Type = models.Admin
		assert.Nil(t, c.CreateUser(user))
	}
	for i := 0; i < 10; i++ {
		user := random.User()
		user.Username = fmt.Sprintf("manager-%s", user.Username)
		user.Type = models.Manager
		assert.Nil(t, c.CreateUser(user))
	}
	t.Run("QueryAll", func(tt *testing.T) {
		results, qErr := c.QueryUsers(&UserFilter{Page: 1})
		assert.Nil(tt, qErr)
		assert.Equal(tt, DefaultPageSize, len(results.Results))
		assert.Equal(tt, 20/DefaultPageSize, int(results.TotalPages))
	})
	t.Run("QueryAdminsByType", func(tt *testing.T) {
		adminType := models.Admin
		results, qErr := c.QueryUsers(&UserFilter{Page: 1, Type: &adminType})
		assert.Nil(tt, qErr)
		assert.Equal(tt, 10, len(results.Results))
		assert.Equal(tt, 10/DefaultPageSize+1, int(results.TotalPages))
	})
	t.Run("QueryManagersByType", func(tt *testing.T) {
		managerType := models.Manager
		results, qErr := c.QueryUsers(&UserFilter{Page: 1, Type: &managerType})
		assert.Nil(tt, qErr)
		assert.Equal(tt, 10, len(results.Results))
		assert.Equal(tt, 10/DefaultPageSize+1, int(results.TotalPages))
	})
}
