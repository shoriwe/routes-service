package models

import (
	"testing"

	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&User{}))
	t.Run("ValidUser", func(tt *testing.T) {
		assert.Nil(tt, db.Create(&User{
			Username: "sulcud",
			Password: "password",
			Type:     Admin,
		}).Error)
	})
	t.Run("RepeatedUser", func(tt *testing.T) {
		assert.Nil(tt, db.Create(&User{
			Username: "sulcud2",
			Password: "password",
			Type:     Admin,
		}).Error)
		assert.NotNil(tt, db.Create(&User{
			Username: "sulcud2",
			Password: "password",
			Type:     Admin,
		}).Error)
	})
}

func TestUser_Claims(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	u := &User{
		Username: "sulcud",
		Password: "password",
		Type:     Admin,
	}
	assert.Nil(t, db.AutoMigrate(&User{}))
	assert.Nil(t, db.Create(u).Error)
	assert.NotNil(t, u.Claims())
}

func TestUser_FromClaims(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	u := &User{
		Username: "sulcud",
		Password: "password",
		Type:     Admin,
	}
	assert.Nil(t, db.AutoMigrate(&User{}))
	assert.Nil(t, db.Create(u).Error)
	t.Run("ValidFromClaims", func(tt *testing.T) {
		var u2 User
		assert.Nil(tt, u2.FromClaims(u.Claims()))
		assert.Equal(tt, u.UUID, u2.UUID)
	})
	t.Run("NoUUID", func(tt *testing.T) {
		var u2 User
		claims := u.Claims()
		delete(claims, "uuid")
		assert.NotNil(tt, u2.FromClaims(claims))
	})
}

func TestUser_Authenticate(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	u := &User{
		Username: "sulcud",
		Password: "password",
		Type:     Admin,
	}
	assert.Nil(t, db.AutoMigrate(&User{}))
	assert.Nil(t, db.Create(u).Error)
	t.Run("ValidPassword", func(tt *testing.T) {
		assert.True(tt, u.Authenticate("password"))
	})
	t.Run("InvalidPassword", func(tt *testing.T) {
		assert.False(tt, u.Authenticate("wrong"))
	})
}

func TestUser_BeforeSave(t *testing.T) {
	db := sqlite.NewTest()
	conn, _ := db.DB()
	defer conn.Close()
	assert.Nil(t, db.AutoMigrate(&User{}))
	t.Run("ValidUser", func(tt *testing.T) {
		assert.Nil(tt, db.Create(&User{
			Username: "sulcud",
			Password: "password",
			Type:     Admin,
		}).Error)
	})
	t.Run("NoPassword", func(tt *testing.T) {
		assert.NotNil(tt, db.Create(&User{
			Username: "sulcud2",
			Type:     Admin,
		}).Error)
	})
}
