package api

const (
	CredentialsKey = "CREDENTIALS_KEY"
)

type Status struct {
	Succeed bool   `json:"succeed"`
	Error   string `json:"error,omitempty"`
}

var (
	StatusOK = Status{
		Succeed: true,
		Error:   "Operation succeed",
	}
	UnauthorizedStatus = Status{
		Succeed: false,
		Error:   "Unauthorized",
	}
	ForbiddenStatus = Status{
		Succeed: false,
		Error:   "Forbidden",
	}
)

const (
	RootRoute = "/api"
)

const (
	UUIDParam = "uuid"
)

/*
- Real time monitoring of vehicle locations
- Login
- User administration
- Export data to JSON
- Calculate routes from A to B
- Download map
*/

/*
Vehicles API keys can:
- [ ] Notify location
- [ ] Plan route

Public
- [X] Download map
- [X] Login

Admin user can:
- [X] CRUD users
- [ ] CRUD vehicles
- [ ] CRUD API keys
- [ ] Monitor location
- [ ] Export data

Manager user can:
- [ ] Monitor location
- [ ] Export data
*/
const (
	EchoRoute              = "/echo"
	MapRoute               = "/map"
	LoginRoute             = "/login"
	UserRoute              = "/user"
	UserRouteWithParams    = UserRoute + "/:" + UUIDParam
	VehicleRoute           = "/vehicle"
	VehicleRouteWithParams = VehicleRoute + "/:" + UUIDParam
	LocationRoute          = "/location"
	ExportRoute            = "/export"
	RoutePlanningRoute     = "/route"
)
