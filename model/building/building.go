package BuildingModel

import (
	db "github.com/taufikardiyan28/rmet1/db"
)

type Building struct {
	BuildID             int64          `json:"build_id" db:"build_id"`
	BuildAddress        db.NullString  `json:"build_address,omitempty" db:"build_address"`
	BuildAddressPostal  db.NullInt64   `json:"build_address_postal,omitempty" db:"build_address_postal"`
	BuildAddressRt      db.NullInt64   `json:"build_address_rt,omitempty" db:"build_address_rt"`
	BuildAddressRw      db.NullInt64   `json:"build_address_rw,omitempty" db:"build_address_rw"`
	BuildAudit          db.NullInt64   `json:"build_audit,omitempty" db:"build_audit"`
	BuildCategoryID     db.NullInt64   `json:"build_category_id,omitempty" db:"build_category_id"`
	BuildContractID     db.NullInt64   `json:"build_contract_id,omitempty" db:"build_contract_id"`
	BuildCreateBy       db.NullString  `json:"build_create_by,omitempty" db:"build_create_by"`
	BuildCreateDate     db.NullTime    `json:"build_create_date,omitempty" db:"build_create_date"`
	BuildDelStatus      db.NullString  `json:"build_del_status,omitempty" db:"build_del_status"`
	BuildDescriptions   db.NullString  `json:"build_descriptions,omitempty" db:"build_descriptions"`
	BuildEndContract    db.NullTime    `json:"build_end_contract,omitempty" db:"build_end_contract"`
	BuildGoogleID       db.NullString  `json:"build_google_id,omitempty" db:"build_google_id"`
	BuildGracePeriodID  db.NullInt64   `json:"build_grace_period_id,omitempty" db:"build_grace_period_id"`
	BuildKabupaten      db.NullString  `json:"build_kabupaten,omitempty" db:"build_kabupaten"`
	BuildKecamatan      db.NullString  `json:"build_kecamatan,omitempty" db:"build_kecamatan"`
	BuildKelurahan      db.NullString  `json:"build_kelurahan,omitempty" db:"build_kelurahan"`
	BuildLatitude       db.NullString  `json:"build_latitude" db:"build_latitude"`
	BuildLongitude      db.NullString  `json:"build_longitude" db:"build_longitude"`
	BuildName           db.NullString  `json:"build_name,omitempty" db:"build_name"`
	BuildNearby         db.NullString  `json:"build_nearby,omitempty" db:"build_nearby"`
	BuildNoContract     db.NullString  `json:"build_no_contract,omitempty" db:"build_no_contract"`
	BuildPhotos         db.NullString  `json:"build_photos,omitempty" db:"build_photos"`
	BuildProdName       db.NullString  `json:"build_prod_name,omitempty" db:"build_prod_name"`
	BuildProvinsi       db.NullString  `json:"build_provinsi,omitempty" db:"build_provinsi"`
	BuildRateCount      db.NullInt64   `json:"build_rate_count" db:"build_rate_count"`
	BuildRateValue      db.NullFloat64 `json:"build_rate_value" db:"build_rate_value"`
	BuildRecomendation  db.NullInt64   `json:"build_recomendation" db:"build_recomendation"`
	BuildService        db.NullString  `json:"build_service,omitempty" db:"build_service"`
	BuildStartContract  db.NullTime    `json:"build_start_contract" db:"build_start_contract"`
	BuildStatusBuilding db.NullInt64   `json:"build_status_building,omitempty" db:"build_status_building"`
	BuildSubProd        db.NullString  `json:"build_sub_prod,omitempty" db:"build_sub_prod"`
	BuildTotalRoom      db.NullString  `json:"build_total_room" db:"build_total_room"`
	BuildUpdateBy       db.NullString  `json:"build_update_by,omitempty" db:"build_update_by"`
	BuildUpdateDate     db.NullTime    `json:"build_update_date" db:"build_update_date"`
}
