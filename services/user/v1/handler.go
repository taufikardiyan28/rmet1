package userv1

import (
	"github.com/labstack/echo/v4"
	user "github.com/taufikardiyan28/rmet1/services/user"
)

type (
	Service struct {
		user.Service
	}
)

func (r *Service) RegisterRouters(g *echo.Group) {
	g.POST("/list", r.List)
}
