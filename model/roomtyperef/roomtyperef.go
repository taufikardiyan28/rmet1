package RoomTypeRefModel

import (
	"time"
)

type RoomTypeRef struct {
	RoomtypeRefID         int64     `json:"roomtype_ref_id" db:"roomtype_ref_id"`
	RoomtypeRefBed        string    `json:"roomtype_ref_bed" db:"roomtype_ref_bed"`
	RoomtypeRefCreateBy   string    `json:"roomtype_ref_create_by" db:"roomtype_ref_create_by"`
	RoomtypeRefCreateDate time.Time `json:"roomtype_ref_create_date" db:"roomtype_ref_create_date"`
	RoomtypeRefDelStatus  string    `json:"roomtype_ref_del_status" db:"roomtype_ref_del_status"`
	RoomtypeRefName       string    `json:"roomtype_ref_name" db:"roomtype_ref_name"`
	RoomtypeRefUpdateBy   string    `json:"roomtype_ref_update_by" db:"roomtype_ref_update_by"`
	RoomtypeRefUpdateDate time.Time `json:"roomtype_ref_update_date" db:"roomtype_ref_update_date"`
}
