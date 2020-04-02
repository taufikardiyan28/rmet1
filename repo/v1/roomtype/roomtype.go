package RoomTypeRepo

import (
	db "github.com/taufikardiyan28/rmet1/db"
	h "github.com/taufikardiyan28/rmet1/helper"
	RoomTypeModel "github.com/taufikardiyan28/rmet1/model/roomtype"
)

type (
	Repo struct {
		db.DB
	}
)

func (r *Repo) getDefaultQuery() string {
	return `SELECT * FROM room_type WHERE roomtype_del_status='0' `
}

func (r *Repo) List(paging *h.PagingData) ([]RoomTypeModel.RoomType, int, error) {
	strSql := r.getDefaultQuery()

	data := []RoomTypeModel.RoomType{}
	period := h.PeriodFilter{
		DateColumn: "roomtype_create_date",
		StartDate:  paging.StartCreateDate,
		EndDate:    paging.EndCreateDate + " 23:59:59",
	}
	totalData, err := r.GetPaging(&data, strSql, paging, period)
	return data, totalData, err
}
