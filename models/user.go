package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	DefaultPasswordCost    = 12
	DefaultExpirationDelta = 30 * 24 * time.Hour
)

type UserType int

const (
	Admin UserType = iota
	Manager
)

type User struct {
	Model
	Username     string   `gorm:"unique;not null;" json:"username"`
	Password     string   `gorm:"-" json:"password"`
	PasswordHash []byte   `gorm:"not null;" json:"-"`
	Type         UserType `gorm:"not null;" json:"account_type"`
}

func (u *User) Claims() jwt.MapClaims {
	return jwt.MapClaims{
		"uuid": u.UUID.String(),
		"exp":  time.Now().Add(DefaultExpirationDelta).Unix(),
	}
}

func (u *User) FromClaims(m jwt.MapClaims) error {
	userUUID, ok := m["uuid"]
	if !ok {
		return fmt.Errorf("incomplete UUID")
	}
	u.UUID = uuid.FromStringOrNil(userUUID.(string))
	return nil
}

func (u *User) Authenticate(password string) bool {
	return u.PasswordHash != nil && bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) == nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	bErr := u.Model.BeforeSave(tx)
	if bErr != nil {
		return bErr
	}
	if len(u.Password) == 0 {
		return fmt.Errorf("password is empty")
	}
	var err error
	u.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(u.Password), DefaultPasswordCost)
	if err != nil {
		return err
	}
	return nil
}
