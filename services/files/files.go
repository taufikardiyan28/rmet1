package files

import (
	"github.com/labstack/echo/v4"
	mid "github.com/taufikardiyan28/rmet1/middleware"
)

type (
	Service struct {
		mid.Request
	}
)

func (s *Service) Download(c echo.Context) error {
	type param struct {
		FileName string `json:"file_name" validate:"required"`
	}

	p := new(param)
	if err := s.BindAndValidate(c, p); err != nil {
		return s.Send(c, 400, "Error", err.Error(), "")
	}
	return c.File("./static/files/" + p.FileName)
}

func RegisterRouters(g *echo.Group) {
	s := Service{}
	g.POST("/get/file", s.Download)
}
