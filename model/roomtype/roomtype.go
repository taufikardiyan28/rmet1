package RoomTypeModel

import (
	"time"
)

type RoomType struct {
	RoomtypeID          int64     `json:"roomtype_id" db:"roomtype_id"`
	RoomtypeBed         string    `json:"roomtype_bed" db:"roomtype_bed"`
	RoomtypeBuildID     int64     `json:"roomtype_build_id" db:"roomtype_build_id"`
	RoomtypeComPrice    int64     `json:"roomtype_com_price" db:"roomtype_com_price"`
	RoomtypeCreateBy    string    `json:"roomtype_create_by" db:"roomtype_create_by"`
	RoomtypeCreateDate  time.Time `json:"roomtype_create_date" db:"roomtype_create_date"`
	RoomtypeDelStatus   string    `json:"roomtype_del_status" db:"roomtype_del_status"`
	RoomtypeDescription string    `json:"roomtype_description" db:"roomtype_description"`
	RoomtypeIsi         int64     `json:"roomtype_isi" db:"roomtype_isi"`
	RoomtypeKm          string    `json:"roomtype_km" db:"roomtype_km"`
	RoomtypeKmTotal     int64     `json:"roomtype_km_total" db:"roomtype_km_total"`
	RoomtypeKosong      int64     `json:"roomtype_kosong" db:"roomtype_kosong"`
	RoomtypeName        string    `json:"roomtype_name" db:"roomtype_name"`
	RoomtypePhotos      string    `json:"roomtype_photos" db:"roomtype_photos"`
	RoomtypePrice       int64     `json:"roomtype_price" db:"roomtype_price"`
	RoomtypePricePromo  int64     `json:"roomtype_price_promo" db:"roomtype_price_promo"`
	RoomtypePromoPrice  int64     `json:"roomtype_promo_price" db:"roomtype_promo_price"`
	RoomtypeRateCount   int64     `json:"roomtype_rate_count" db:"roomtype_rate_count"`
	RoomtypeRateValue   float64   `json:"roomtype_rate_value" db:"roomtype_rate_value"`
	RoomtypeSize        string    `json:"roomtype_size" db:"roomtype_size"`
	RoomtypeTotal       int64     `json:"roomtype_total" db:"roomtype_total"`
	RoomtypeUpdateBy    string    `json:"roomtype_update_by" db:"roomtype_update_by"`
	RoomtypeUpdateDate  time.Time `json:"roomtype_update_date" db:"roomtype_update_date"`
}
