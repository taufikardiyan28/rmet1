package userv1

import (
	"github.com/labstack/echo/v4"
	roomtype "github.com/taufikardiyan28/rmet1/services/roomtype"
)

type (
	Service struct {
		roomtype.Service
	}
)

func (r *Service) RegisterRouters(g *echo.Group) {
	g.POST("/list", r.List)
}
