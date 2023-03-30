package api

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/routes-service/common/random"
	"github.com/shoriwe/routes-service/common/sqlite"
	"github.com/shoriwe/routes-service/controller"
	"github.com/shoriwe/routes-service/maps/samples"
)

type Handler struct {
	Controller *controller.Controller
}

func New(c *controller.Controller) *gin.Engine {
	h := Handler{Controller: c}
	engine := gin.Default()
	root := engine.Group(RootRoute)
	// - Public
	root.GET(MapRoute, h.Map)
	root.POST(LoginRoute, h.Login)
	// User auth required
	user := root.Group("/", h.checkJWT)
	user.GET(EchoRoute, h.Echo)
	// - All users
	// -- Monitor Locations
	// -- Export
	// - Admin
	admin := user.Group("/", h.onlyAdmin)
	// -- User's CRUD
	admin.PUT(UserRoute, h.CreateUser)
	admin.DELETE(UserRouteWithParams, h.DeleteUser)
	admin.PATCH(UserRoute, h.UpdateUser)
	admin.POST(UserRoute, h.QueryUsers)
	// -- Vehicle's CRUD
	admin.PUT(VehicleRoute, h.CreateVehicle)
	admin.DELETE(VehicleRouteWithParams, h.DeleteVehicle)
	admin.PATCH(VehicleRoute, h.UpdateVehicle)
	admin.POST(VehicleRoute, h.QueryVehicles)
	// -- API's CRUD
	// - Vehicles
	return engine
}

func NewTest(t *testing.T) (*controller.Controller, *httpexpect.Expect, func()) {
	c := controller.New([]byte(random.String()), sqlite.NewTest(), &samples.ImaginaryCity)
	engine := New(c)
	server := httptest.NewServer(engine)
	return c, httpexpect.Default(t, server.URL+RootRoute), server.Close
}
