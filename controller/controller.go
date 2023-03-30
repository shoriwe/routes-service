package controller

import (
	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/common/session"
	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/shoriwe/routes-service/maps"
	"github.com/shoriwe/routes-service/maps/samples"
	"github.com/shoriwe/routes-service/models"
	"gorm.io/gorm"
)

type Controller struct {
	Session *session.JWT
	DB      *gorm.DB
	Map     *maps.Map
}

func (c *Controller) Close() {
	conn, err := c.DB.DB()
	if err != nil {
		panic(err)
	}
	err = conn.Close()
	if err != nil {
		panic(err)
	}
}

func New(secret []byte, db *gorm.DB, m *maps.Map) *Controller {
	c := &Controller{
		Session: session.New(secret),
		DB:      db,
		Map:     m,
	}
	c.DB.AutoMigrate(
		&models.User{}, &models.Vehicle{},
		&models.Route{}, &models.Location{},
		&models.APIKey{},
	)
	return c
}

func NewTest() *Controller {
	return New([]byte(random.String()), sqlite.NewTest(), &samples.ImaginaryCity)
}
