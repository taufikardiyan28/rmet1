package OwnerModel

import (
	db "github.com/taufikardiyan28/rmet1/db"
)

type Owner struct {
	OwnerID            int64         `db:"owner_id" json:"owner_id"`
	OwnerAddress       db.NullString `db:"owner_address" json:"owner_address"`
	OwnerBankName      db.NullString `json:"owner_bank_name" db:"owner_bank_name"`
	OwnerBuildingID    db.NullInt64  `json:"owner_building_id" db:"owner_building_id"`
	OwnerDelStatus     db.NullString `json:"owner_del_status" db:"owner_del_status"`
	OwnerDob           db.NullTime   `json:"owner_dob" db:"owner_dob"`
	OwnerEmail         db.NullString `json:"owner_email" db:"owner_email"`
	OwnerFirstnameBc   db.NullString `json:"owner_firstname_bc" db:"owner_firstname_bc"`
	OwnerLastnameBc    db.NullString `json:"owner_lastname_bc" db:"owner_lastname_bc"`
	OwnerMaritalStatus db.NullString `json:"owner_marital_status" db:"owner_marital_status"`
	OwnerName          db.NullString `json:"owner_name" db:"owner_name"`
	OwnerNik           db.NullString `json:"owner_nik" db:"owner_nik"`
	OwnerNoRek         db.NullString `json:"owner_no_rek" db:"owner_no_rek"`
	OwnerPayMetode     db.NullString `json:"owner_pay_metode" db:"owner_pay_metode"`
	OwnerPhone         db.NullString `json:"owner_phone" db:"owner_phone"`
	OwnerPriceContract db.NullInt64  `json:"owner_price_contract" db:"owner_price_contract"`
	OwnerRekName       db.NullString `json:"owner_rek_name" db:"owner_rek_name"`
	OwnerCreateBy      db.NullString `json:"owner_create_by" db:"owner_create_by"`
	OwnerCreateDate    db.NullTime   `json:"owner_create_date" db:"owner_create_date"`
	OwnerUpdateBy      db.NullString `json:"owner_update_by" db:"owner_update_by"`
	OwnerUpdateDate    db.NullTime   `json:"owner_update_date" db:"owner_update_date"`
}
