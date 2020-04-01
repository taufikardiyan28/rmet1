package ProvinsiModel

import (
	"time"
)

type Provinsi struct {
	ProvID         int       `json:"prov_id" db:"prov_id"`
	ProvCreateDate time.Time `json:"prov_create_date" db:"prov_create_date"`
	ProvName       string    `json:"prov_name" db:"prov_name"`
}
