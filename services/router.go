package router

import (
	"github.com/labstack/echo/v4"
	acquisitionv1 "github.com/taufikardiyan28/rmet1/services/acquisition/v1"
	roomtypev1 "github.com/taufikardiyan28/rmet1/services/roomtype/v1"
	userv1 "github.com/taufikardiyan28/rmet1/services/user/v1"
)

func RegisterV1(g *echo.Group) {
	f := g.Group("/form")
	{
		//v1/form/acquisition
		acqV1 := acquisitionv1.Service{}
		acqV1.RegisterRouters(f.Group("/acquisition"))
	}

	//v1/user
	userV1 := userv1.Service{}
	userV1.RegisterRouters(g.Group("/user"))

	//v1/roomtype
	rmtypeV1 := roomtypev1.Service{}
	rmtypeV1.RegisterRouters(g.Group("/roomtype"))
}
