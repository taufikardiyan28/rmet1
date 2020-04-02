package buildingv1

import (
	"github.com/labstack/echo/v4"
	building "github.com/taufikardiyan28/rmet1/services/acquisition"
)

type (
	Service struct {
		building.Service
	}
)

func (r *Service) RegisterRouters(g *echo.Group) {
	g.POST("/list", r.List)
	g.POST("/summary", r.Summary)
	g.POST("/detail", r.Detail)
	g.POST("/find-place-id", r.GetGooglePlaceID)
	g.POST("/get-latlng", r.GetLatLng)
	g.PUT("/updateroom", r.UpdateRoom)
}
