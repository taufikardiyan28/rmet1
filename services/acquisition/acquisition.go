package acquisition

import (
	"math"
	"sync"

	"github.com/labstack/echo/v4"
	h "github.com/taufikardiyan28/rmet1/helper"
	mid "github.com/taufikardiyan28/rmet1/middleware"
	repo "github.com/taufikardiyan28/rmet1/repo/v1/acquisition"
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

func (s *Service) Summary(c echo.Context) error {
	rep := new(repo.Repo)
	data, err := rep.Summary()
	if err != nil {
		return s.Send(c, 400, "Error", err.Error(), data)
	}

	return s.Send(c, 200, "Success", "Success", data)
}

func (s *Service) Detail(c echo.Context) error {
	type param struct {
		BuildId int `json:"build_id"`
	}

	p := new(param)
	if err := s.BindAndValidate(c, p); err != nil {
		return s.Send(c, 400, "Error", err.Error(), "")
	}

	detail := make(chan interface{}, 1)
	doc := make(chan interface{}, 1)
	roomtype := make(chan interface{}, 1)
	auditlog := make(chan interface{}, 1)
	errors := make(chan error, 2)
	rep := new(repo.Repo)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		data, err := rep.GetDetail(p.BuildId)
		if err != nil {
			errors <- err
			return
		}
		detail <- data
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		data, err := rep.GetDocument(p.BuildId)
		if err != nil {
			errors <- err
			return
		}
		doc <- data
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		data, err := rep.GetRoomType(p.BuildId)
		if err != nil {
			errors <- err
			return
		}
		roomtype <- data
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		data, err := rep.GetAuditLog(p.BuildId)
		if err != nil {
			errors <- err
			return
		}
		auditlog <- data
	}()

	go func() {
		wg.Wait()
		close(detail)
		close(doc)
		close(roomtype)
		close(auditlog)
		close(errors)
	}()

	for err := range errors {
		return s.Send(c, 400, "Error", err.Error(), nil)
	}

	resp := make(map[string]interface{})
	for res := range detail {
		resp["unitinfo"] = res
	}

	for res := range doc {
		resp["document"] = res
	}

	for res := range roomtype {
		resp["roomtype"] = res
	}

	for res := range auditlog {
		resp["auditlog"] = res
	}

	return s.Send(c, 200, "Success", "Success", map[string]interface{}{"unitinfo": resp})
}

func (s *Service) GetGooglePlaceID(c echo.Context) error {
	type param struct {
		Search string `json:"search" validate:"required"`
	}

	p := new(param)
	if err := s.BindAndValidate(c, p); err != nil {
		return s.Send(c, 400, "Error", err.Error(), "")
	}

	rep := new(repo.Repo)

	placeId, err := rep.GetGooglePlaceId(p.Search)
	if err != nil {
		return s.Send(c, 400, "Error", err.Error(), "")
	}
	return s.Send(c, 200, "Success", "Success", map[string]interface{}{"place_id": placeId})
}

func (s *Service) GetLatLng(c echo.Context) error {
	type param struct {
		PlaceID string `json:"place_id" validate:"required"`
	}

	p := new(param)
	if err := s.BindAndValidate(c, p); err != nil {
		return s.Send(c, 400, "Error", err.Error(), "")
	}

	rep := new(repo.Repo)

	lat, lng, err := rep.GetLatLng(p.PlaceID)
	if err != nil {
		return s.Send(c, 400, "Error", err.Error(), "")
	}
	return s.Send(c, 200, "Success", "Success", map[string]interface{}{"lat": lat, "lng": lng})
}
