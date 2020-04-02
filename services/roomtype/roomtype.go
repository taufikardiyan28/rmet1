package roomtype

import (
	"math"

	"github.com/labstack/echo/v4"
	h "github.com/taufikardiyan28/rmet1/helper"
	mid "github.com/taufikardiyan28/rmet1/middleware"
	repo "github.com/taufikardiyan28/rmet1/repo/v1/roomtype"
)

type (
	Service struct {
		mid.Request
	}
)

func (s *Service) List(c echo.Context) error {
	paging := new(h.PagingData)
	if err := s.BindAndValidate(c, paging); err != nil {
		return s.Send(c, 400, "Error", err.Error(), "")
	}

	rep := new(repo.Repo)
	data, totalData, err := rep.List(paging)
	if err != nil {
		return s.Send(c, 400, "Error", err.Error(), data)
	}

	totalPage := math.Ceil(float64(totalData) / float64(paging.Show))
	respData := mid.PagingResponse{
		ListResult: data,
		Page:       paging.Page,
		TotalData:  totalData,
		TotalPage:  int(totalPage),
	}
	return s.Send(c, 200, "Success", "Success", respData)
}
