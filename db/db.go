package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	tx *sqlx.Tx
}

var Pool *sqlx.DB

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
