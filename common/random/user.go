package random

import (
	"github.com/shoriwe/routes-service/models"
)

func User() *models.User {
	return &models.User{
		Username: String(),
		Password: String()[:72],
	}
}
