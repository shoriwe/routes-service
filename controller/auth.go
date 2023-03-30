package controller

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shoriwe/routes-service/models"
)

func (c *Controller) Login(credentials *models.User) (string, error) {
	user := &models.User{}
	fErr := c.DB.Where("username = ? AND deleted_at IS NULL", credentials.Username).First(user).Error
	if fErr != nil {
		return "", fErr
	}
	if !user.Authenticate(credentials.Password) {
		return "", ErrorUnauthorized
	}
	return c.Session.New(user.Claims()), nil
}

func (c *Controller) UserValid(user *models.User) (*models.User, error) {
	credentials := &models.User{}
	err := c.DB.First(credentials, "uuid = ? AND deleted_at IS NULL", user.UUID).Error
	if err != nil {
		return nil, fmt.Errorf("login error: %w", err)
	}
	return credentials, nil
}

func (c *Controller) AuthorizeUser(tokenString string) (*models.User, error) {
	token, tErr := jwt.Parse(tokenString, c.Session.KeyFunc)
	if tErr != nil {
		return nil, tErr
	}
	jwtUser := &models.User{}
	fErr := jwtUser.FromClaims(token.Claims.(jwt.MapClaims))
	if fErr != nil {
		return nil, fErr
	}
	credentials, aErr := c.UserValid(jwtUser)
	if aErr != nil {
		return nil, aErr
	}
	return credentials, nil
}
