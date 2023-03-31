package api

import (
	"net/http/httptest"
	"net/url"
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
	root := engine.Group(APIRoute)
	// - Public
	root.GET(MapRoute, h.Map)
	root.POST(LoginRoute, h.Login)
	// User auth required
	user := root.Group(RootRoute, h.CheckJWT)
	// - All users
	// -- Echo
	user.GET(EchoRoute, h.Echo)
	// -- Monitor Locations
	root.GET(LocationConsumerRoute, h.LocationListener)
	// -- Export
	// - Admin
	admin := user.Group(RootRoute, h.OnlyAdmin)
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
	admin.PUT(APIKeyRoute, h.CreateAPIKey)
	admin.DELETE(APIKeyRouteWithParams, h.DeleteAPIKey)
	admin.POST(APIKeyRoute, h.QueryAPIKeys)
	// - Vehicles
	// vehicles := root.Group(RootRoute, h.CheckAPIKey)
	root.GET(LocationProducerRoute, h.LocationProducer)
	return engine
}

func NewTest(t *testing.T) (*controller.Controller, *httpexpect.Expect, func()) {
	c := controller.New([]byte(random.String()), sqlite.NewTest(), &samples.ImaginaryCity)
	engine := New(c)
	server := httptest.NewServer(engine)
	return c, httpexpect.Default(t, server.URL+APIRoute), server.Close
}

func NewTestWS(t *testing.T) (*controller.Controller, string, func()) {
	c := controller.New([]byte(random.String()), sqlite.NewTest(), &samples.ImaginaryCity)
	engine := New(c)
	server := httptest.NewServer(engine)
	u, _ := url.Parse(server.URL + APIRoute)
	u.Scheme = "ws"
	return c, u.String(), server.Close
}
