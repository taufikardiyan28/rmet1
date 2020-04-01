package BuildingModel

import (
	"time"
)

type Building struct {
	BuildID             int64     `json:"build_id" db:"build_id"`
	BuildAddress        string    `json:"build_address" db:"build_address"`
	BuildAddressPostal  int64     `json:"build_address_postal" db:"build_address_postal"`
	BuildAddressRt      int64     `json:"build_address_rt" db:"build_address_rt"`
	BuildAddressRw      int64     `json:"build_address_rw" db:"build_address_rw"`
	BuildAudit          int64     `json:"build_audit" db:"build_audit"`
	BuildCategoryID     int64     `json:"build_category_id" db:"build_category_id"`
	BuildContractID     int64     `json:"build_contract_id" db:"build_contract_id"`
	BuildCreateBy       string    `json:"build_create_by" db:"build_create_by"`
	BuildCreateDate     time.Time `json:"build_create_date" db:"build_create_date"`
	BuildDelStatus      string    `json:"build_del_status" db:"build_del_status"`
	BuildDescriptions   string    `json:"build_descriptions" db:"build_descriptions"`
	BuildEndContract    time.Time `json:"build_end_contract" db:"build_end_contract"`
	BuildGoogleID       string    `json:"build_google_id" db:"build_google_id"`
	BuildGracePeriodID  int64     `json:"build_grace_period_id" db:"build_grace_period_id"`
	BuildKabupaten      string    `json:"build_kabupaten" db:"build_kabupaten"`
	BuildKecamatan      string    `json:"build_kecamatan" db:"build_kecamatan"`
	BuildKelurahan      string    `json:"build_kelurahan" db:"build_kelurahan"`
	BuildLatitude       string    `json:"build_latitude" db:"build_latitude"`
	BuildLongitude      string    `json:"build_longitude" db:"build_longitude"`
	BuildName           string    `json:"build_name" db:"build_name"`
	BuildNearby         string    `json:"build_nearby" db:"build_nearby"`
	BuildNoContract     string    `json:"build_no_contract" db:"build_no_contract"`
	BuildPhotos         string    `json:"build_photos" db:"build_photos"`
	BuildProdName       string    `json:"build_prod_name" db:"build_prod_name"`
	BuildProvinsi       string    `json:"build_provinsi" db:"build_provinsi"`
	BuildRateCount      int64     `json:"build_rate_count" db:"build_rate_count"`
	BuildRateValue      float64   `json:"build_rate_value" db:"build_rate_value"`
	BuildRecomendation  int64     `json:"build_recomendation" db:"build_recomendation"`
	BuildService        string    `json:"build_service" db:"build_service"`
	BuildStartContract  time.Time `json:"build_start_contract" db:"build_start_contract"`
	BuildStatusBuilding int64     `json:"build_status_building" db:"build_status_building"`
	BuildSubProd        string    `json:"build_sub_prod" db:"build_sub_prod"`
	BuildTotalRoom      string    `json:"build_total_room" db:"build_total_room"`
	BuildUpdateBy       string    `json:"build_update_by" db:"build_update_by"`
	BuildUpdateDate     time.Time `json:"build_update_date" db:"build_update_date"`
}
