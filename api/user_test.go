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

func TestHandler_CreateUser(t *testing.T) {
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
		expect.PUT(UserRoute).WithHeader("Authorization", adminToken).WithJSON(random.User()).Expect().Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		expect.PUT(UserRoute).WithHeader("Authorization", managerToken).WithJSON(random.User()).Expect().Status(http.StatusForbidden)
	})
	t.Run("NoJSON", func(tt *testing.T) {
		expect.PUT(UserRoute).
			WithHeader("Authorization", adminToken).
			WithHeader("Content-Type", "application/json").
			WithText("{").
			Expect().
			Status(http.StatusBadRequest)
	})
	t.Run("RepeatedUser", func(tt *testing.T) {
		user := random.User()
		expect.PUT(UserRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(user).
			Expect().
			Status(http.StatusOK)
		expect.PUT(UserRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(user).
			Expect().
			Status(http.StatusInternalServerError)
	})
}

func TestHandler_DeleteUser(t *testing.T) {
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
		user := random.User()
		c.CreateUser(user)
		expect.DELETE(UserRoute+fmt.Sprintf("/%s", user.UUID)).
			WithHeader("Authorization", adminToken).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		user := random.User()
		c.CreateUser(user)
		expect.DELETE(UserRoute+fmt.Sprintf("/%s", user.UUID)).
			WithHeader("Authorization", managerToken).
			Expect().
			Status(http.StatusForbidden)
	})
}

func TestHandler_UpdateUser(t *testing.T) {
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
		user := random.User()
		user.Type = models.Manager
		c.CreateUser(user)
		user.Password = ""
		user.Type = models.Admin
		expect.PATCH(UserRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(user).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		user := random.User()
		user.Type = models.Manager
		c.CreateUser(user)
		user.Password = ""
		user.Type = models.Admin
		expect.PATCH(UserRoute).
			WithHeader("Authorization", managerToken).
			WithJSON(user).
			Expect().
			Status(http.StatusForbidden)
	})
	t.Run("InvalidJSON", func(tt *testing.T) {
		user := random.User()
		user.Type = models.Manager
		c.CreateUser(user)
		user.Password = ""
		user.Type = models.Admin
		expect.PATCH(UserRoute).
			WithHeader("Authorization", adminToken).
			WithText("{").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
}

func TestHandler_QueryUsers(t *testing.T) {
	c, expect, serverClose := NewTest(t)
	defer serverClose()
	admin := random.User()
	admin.Type = models.Admin
	assert.Nil(t, c.CreateUser(admin))
	manager := random.User()
	manager.Type = models.Manager
	assert.Nil(t, c.CreateUser(manager))
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
	// Tokens
	adminToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(admin.Claims()))))
	managerToken := fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(c.Session.New(manager.Claims()))))
	t.Run("AsAdmin", func(tt *testing.T) {
		filter := &controller.UserFilter{Page: 1}
		expect.POST(UserRoute).
			WithHeader("Authorization", adminToken).
			WithJSON(filter).
			Expect().
			Status(http.StatusOK)
	})
	t.Run("AsManager", func(tt *testing.T) {
		filter := &controller.UserFilter{Page: 1}
		expect.POST(UserRoute).
			WithHeader("Authorization", managerToken).
			WithJSON(filter).
			Expect().
			Status(http.StatusForbidden)
	})
	t.Run("InvalidJSON", func(tt *testing.T) {
		expect.POST(UserRoute).
			WithHeader("Authorization", adminToken).
			WithText("{").
			WithHeader("Content-Type", "application/json").
			Expect().
			Status(http.StatusBadRequest)
	})
}
