package RoomTypeModel

import (
	db "github.com/taufikardiyan28/rmet1/db"
)

type RoomType struct {
	RoomtypeID          int64          `json:"roomtype_id" db:"roomtype_id"`
	RoomtypeBed         db.NullString  `json:"roomtype_bed" db:"roomtype_bed"`
	RoomtypeBuildID     db.NullInt64   `json:"roomtype_build_id" db:"roomtype_build_id"`
	RoomtypeComPrice    db.NullInt64   `json:"roomtype_com_price" db:"roomtype_com_price"`
	RoomtypeCreateBy    db.NullString  `json:"roomtype_create_by" db:"roomtype_create_by"`
	RoomtypeCreateDate  db.NullTime    `json:"roomtype_create_date" db:"roomtype_create_date"`
	RoomtypeDelStatus   db.NullString  `json:"roomtype_del_status" db:"roomtype_del_status"`
	RoomtypeDescription db.NullString  `json:"roomtype_description" db:"roomtype_description"`
	RoomtypeIsi         db.NullInt64   `json:"roomtype_isi" db:"roomtype_isi"`
	RoomtypeKm          db.NullString  `json:"roomtype_km" db:"roomtype_km"`
	RoomtypeKmTotal     db.NullInt64   `json:"roomtype_km_total" db:"roomtype_km_total"`
	RoomtypeKosong      db.NullInt64   `json:"roomtype_kosong" db:"roomtype_kosong"`
	RoomtypeName        db.NullString  `json:"roomtype_name" db:"roomtype_name"`
	RoomtypePhotos      db.NullString  `json:"roomtype_photos" db:"roomtype_photos"`
	RoomtypePrice       db.NullInt64   `json:"roomtype_price" db:"roomtype_price"`
	RoomtypePricePromo  db.NullInt64   `json:"roomtype_price_promo" db:"roomtype_price_promo"`
	RoomtypePromoPrice  db.NullInt64   `json:"roomtype_promo_price" db:"roomtype_promo_price"`
	RoomtypeRateCount   db.NullInt64   `json:"roomtype_rate_count" db:"roomtype_rate_count"`
	RoomtypeRateValue   db.NullFloat64 `json:"roomtype_rate_value" db:"roomtype_rate_value"`
	RoomtypeSize        db.NullString  `json:"roomtype_size" db:"roomtype_size"`
	RoomtypeTotal       db.NullInt64   `json:"roomtype_total" db:"roomtype_total"`
	RoomtypeUpdateBy    db.NullString  `json:"roomtype_update_by" db:"roomtype_update_by"`
	RoomtypeUpdateDate  db.NullTime    `json:"roomtype_update_date" db:"roomtype_update_date"`
}
