package OwnerModel

import "time"

type Owner struct {
	OwnerID            int64     `db:"owner_id" json:"owner_id"`
	OwnerAddress       string    `db:"owner_address" json:"owner_address"`
	OwnerBankName      string    `json:"owner_bank_name" db:"owner_bank_name"`
	OwnerBuildingID    int64     `json:"owner_building_id" db:"owner_building_id"`
	OwnerCreateBy      string    `json:"owner_create_by" db:"owner_create_by"`
	OwnerCreateDate    time.Time `json:"owner_create_date" db:"owner_create_date"`
	OwnerDelStatus     string    `json:"owner_del_status" db:"owner_del_status"`
	OwnerDob           time.Time `json:"owner_dob" db:"owner_dob"`
	OwnerEmail         string    `json:"owner_email" db:"owner_email"`
	OwnerFirstnameBc   string    `json:"owner_firstname_bc" db:"owner_firstname_bc"`
	OwnerLastnameBc    string    `json:"owner_lastname_bc" db:"owner_lastname_bc"`
	OwnerMaritalStatus string    `json:"owner_marital_status" db:"owner_marital_status"`
	OwnerName          string    `json:"owner_name" db:"owner_name"`
	OwnerNik           string    `json:"owner_nik" db:"owner_nik"`
	OwnerNoRek         string    `json:"owner_no_rek" db:"owner_no_rek"`
	OwnerPayMetode     string    `json:"owner_pay_metode" db:"owner_pay_metode"`
	OwnerPhone         string    `json:"owner_phone" db:"owner_phone"`
	OwnerPriceContract int64     `json:"owner_price_contract" db:"owner_price_contract"`
	OwnerRekName       string    `json:"owner_rek_name" db:"owner_rek_name"`
	OwnerUpdateBy      string    `json:"owner_update_by" db:"owner_update_by"`
	OwnerUpdateDate    time.Time `json:"owner_update_date" db:"owner_update_date"`
}
