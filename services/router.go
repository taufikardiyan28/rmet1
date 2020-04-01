package router

import (
	"github.com/labstack/echo/v4"
	acquisitionv1 "github.com/taufikardiyan28/rmet1/services/acquisition/v1"
)

func RegisterV1(g *echo.Group) {
	f := g.Group("/form")
	{
		//v1/form/acquisition
		acqV1 := acquisitionv1.Service{}
		acqV1.RegisterRouters(f.Group("/acquisition"))
	}
}
