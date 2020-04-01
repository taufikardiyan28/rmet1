package BuildStafModel

import (
	db "github.com/taufikardiyan28/rmet1/db"
)

type BuildStaff struct {
	BuildStaffID         int64         `json:"build_staff_id" db:"build_staff_id"`
	BuildStaffAddress    db.NullString `json:"build_staff_address" db:"build_staff_address"`
	BuildStaffBuildID    db.NullInt64  `json:"build_staff_build_id" db:"build_staff_build_id"`
	BuildStaffCreateBy   db.NullString `json:"build_staff_create_by" db:"build_staff_create_by"`
	BuildStaffCreateDate db.NullTime   `json:"build_staff_create_date" db:"build_staff_create_date"`
	BuildStaffDelStatus  db.NullString `json:"build_staff_del_status" db:"build_staff_del_status"`
	BuildStaffDob        db.NullTime   `json:"build_staff_dob" db:"build_staff_dob"`
	BuildStaffEmail      db.NullString `json:"build_staff_email" db:"build_staff_email"`
	BuildStaffGender     db.NullString `json:"build_staff_gender" db:"build_staff_gender"`
	BuildStaffName       db.NullString `json:"build_staff_name" db:"build_staff_name"`
	BuildStaffNik        db.NullString `json:"build_staff_nik" db:"build_staff_nik"`
	BuildStaffPhone      db.NullString `json:"build_staff_phone" db:"build_staff_phone"`
	BuildStaffPin        db.NullString `json:"build_staff_pin" db:"build_staff_pin"`
	BuildStaffToken      db.NullString `json:"build_staff_token" db:"build_staff_token"`
	BuildStaffUpdateBy   db.NullString `json:"build_staff_update_by" db:"build_staff_update_by"`
	BuildStaffUpdateDate db.NullTime   `json:"build_staff_update_date" db:"build_staff_update_date"`
}
