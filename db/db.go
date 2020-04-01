package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	h "github.com/taufikardiyan28/rmet1/helper"
)

type (
	DB struct {
		tx *sqlx.Tx
	}
	NullInt64 struct {
		sql.NullInt64
	}

	NullFloat64 struct {
		sql.NullFloat64
	}

	NullString struct {
		sql.NullString
	}
	NullTime struct {
		mysql.NullTime
	}
)

var Pool *sqlx.DB

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

func InitDB(dsn string) error {
	var err error
	Pool, err = sqlx.Open("mysql", dsn)

	return err
}

func (d *DB) GetPool() *sqlx.DB {
	return Pool
}

func (d *DB) BeginTx() error {
	var err error
	d.tx, err = d.GetPool().Beginx()
	return err
}

func (d *DB) SetTx(tx *sqlx.Tx) {
	d.tx = tx
}

func (d *DB) GetTx() *sqlx.Tx {
	return d.tx
}

func (d *DB) ClearTx() {
	d.tx = nil
}

func (d *DB) CommitTx() error {
	err := d.tx.Commit()
	d.ClearTx()
	return err
}

func (d *DB) RollbackTx() error {
	err := d.tx.Rollback()
	d.ClearTx()
	return err
}

func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	var res sql.Result
	var err error
	if d.tx != nil {
		res, err = d.tx.Exec(query, args...)
	} else {
		res, err = d.GetPool().Exec(query, args...)
	}
	return res, err
}

func (d *DB) Select(dest interface{}, query string, args ...interface{}) error {
	var err error
	if d.tx != nil {
		err = d.tx.Select(dest, query, args...)
	} else {
		err = d.GetPool().Select(dest, query, args...)
	}
	return err
}

func (d *DB) Get(dest interface{}, query string, args ...interface{}) error {
	var err error
	if d.tx != nil {
		err = d.tx.Get(dest, query, args...)
	} else {
		err = d.GetPool().Get(dest, query, args...)
	}
	return err
}

func (d *DB) GetPaging(dest interface{}, query string, paging *h.PagingData, period h.PeriodFilter) (int, error) {
	filter := ""
	fields := []string{}
	var args []interface{}
	var filterRules []h.Search
	if paging.Search != "" {
		err := json.Unmarshal([]byte(paging.Search), &filterRules)
		if err != nil {
			return 0, err
		}
	}

	for _, el := range filterRules {
		switch el.Op {
		case "le":
			fields = append(fields, el.Column+" < ? ")
			args = append(args, el.Value)
		case "ge":
			fields = append(fields, el.Column+" > ? ")
			args = append(args, el.Value)
		case "eq":
			fields = append(fields, el.Column+" = ? ")
			args = append(args, el.Value)
		default:
			fields = append(fields, el.Column+" like ? ")
			args = append(args, "%"+el.Value+"%")
		}
	}

	if period.StartDate != "" {
		fields = append(fields, period.DateColumn+" >= ?")
		fields = append(fields, period.DateColumn+" <= ?")
		args = append(args, period.StartDate)
		args = append(args, period.EndDate)
	}

	page := paging.Page
	if page > 0 {
		page -= 1
	} else {
		page = 0
	}
	if len(fields) > 0 {
		filter = " HAVING " + strings.Join(fields, " AND ")
		query += filter
	}
	sqlCount := `SELECT COUNT(*) FROM (` + query + ") tb;"

	var cnt int
	err := d.GetPool().Get(&cnt, sqlCount, args...)

	if err != nil {
		return cnt, err
	}

	query += ` LIMIT ?,?`

	args = append(args, page*paging.Show)
	args = append(args, paging.Show)

	err = d.GetPool().Select(dest, query, args...)

	return cnt, err
}

func (d *DB) NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func (d *DB) NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func (d *DB) NewNullDateString(str string) string {
	if str == "" {
		return "0000-00-00"
	}
	return str
}
