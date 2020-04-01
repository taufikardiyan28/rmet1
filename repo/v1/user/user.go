package UserRepo

import (
	db "github.com/taufikardiyan28/rmet1/db"
	h "github.com/taufikardiyan28/rmet1/helper"
	UserModel "github.com/taufikardiyan28/rmet1/model/user"
)

type (
	Repo struct {
		db.DB
	}
)

func (r *Repo) getDefaultQuery() string {
	return `SELECT user_id, user_firstname, user_lastname FROM ` + "`user`" + ` WHERE user_del_status='0' `
}

func (r *Repo) List(paging *h.PagingData) ([]UserModel.User, int, error) {
	strSql := r.getDefaultQuery()

	data := []UserModel.User{}
	period := h.PeriodFilter{
		DateColumn: "user_create_date",
		StartDate:  paging.StartCreateDate,
		EndDate:    paging.EndCreateDate + " 23:59:59",
	}
	totalData, err := r.GetPaging(&data, strSql, paging, period)
	return data, totalData, err
}
