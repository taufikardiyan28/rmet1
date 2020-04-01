package AcquisitionRepo

import (
	db "github.com/taufikardiyan28/rmet1/db"
	h "github.com/taufikardiyan28/rmet1/helper"
	BuildingModel "github.com/taufikardiyan28/rmet1/model/building"
)

type (
	Repo struct {
		db.DB
	}

	BuildingData struct {
		BuildingModel.Building
		AuditStatus string `db:"audit_status" json:"audit_status"`
	}

	SummaryData struct {
		AuditStatus string `db:"audit_status" json:"audit_status"`
		Total       int    `db:"total" json:"total"`
	}
)

func (r *Repo) getDefaultQuery() string {
	return `SELECT * FROM (SELECT *, CASE  WHEN build_audit=0 THEN 'Belum Audit' WHEN build_total_room>build_audit THEN 
			'Audit Ulang' ELSE 'Completed' END AS audit_status FROM building) as t `
}

func (r *Repo) getSummaryQuery() string {
	return `SELECT audit_status, COUNT(*) AS total FROM (SELECT *, CASE  WHEN build_audit=0 THEN 'Belum Audit' WHEN build_total_room>build_audit THEN 
			'Audit Ulang' ELSE 'Completed' END AS audit_status FROM building) AS t
			GROUP BY audit_status
			UNION 
			SELECT 'All' AS audit_status, COUNT(*) FROM building`
}

func (r *Repo) List(paging *h.PagingData) ([]BuildingData, int, error) {
	strSql := r.getDefaultQuery()

	data := []BuildingData{}
	period := h.PeriodFilter{
		DateColumn: "build_create_date",
		StartDate:  paging.StartCreateDate,
		EndDate:    paging.EndCreateDate + " 23:59:59",
	}
	totalData, err := r.GetPaging(&data, strSql, paging, period)
	return data, totalData, err
}

func (r *Repo) Summary() ([]SummaryData, error) {
	strSql := r.getSummaryQuery()

	data := []SummaryData{}

	err := r.Select(&data, strSql)
	return data, err
}
