package BuildStafModel

import (
	"time"
)

type BuildStaff struct {
	BuildStaffID         int64     `json:"build_staff_id" db:"build_staff_id"`
	BuildStaffAddress    string    `json:"build_staff_address" db:"build_staff_address"`
	BuildStaffBuildID    int64     `json:"build_staff_build_id" db:"build_staff_build_id"`
	BuildStaffCreateBy   string    `json:"build_staff_create_by" db:"build_staff_create_by"`
	BuildStaffCreateDate time.Time `json:"build_staff_create_date" db:"build_staff_create_date"`
	BuildStaffDelStatus  string    `json:"build_staff_del_status" db:"build_staff_del_status"`
	BuildStaffDob        time.Time `json:"build_staff_dob" db:"build_staff_dob"`
	BuildStaffEmail      string    `json:"build_staff_email" db:"build_staff_email"`
	BuildStaffGender     string    `json:"build_staff_gender" db:"build_staff_gender"`
	BuildStaffName       string    `json:"build_staff_name" db:"build_staff_name"`
	BuildStaffNik        string    `json:"build_staff_nik" db:"build_staff_nik"`
	BuildStaffPhone      string    `json:"build_staff_phone" db:"build_staff_phone"`
	BuildStaffPin        string    `json:"build_staff_pin" db:"build_staff_pin"`
	BuildStaffToken      string    `json:"build_staff_token" db:"build_staff_token"`
	BuildStaffUpdateBy   string    `json:"build_staff_update_by" db:"build_staff_update_by"`
	BuildStaffUpdateDate time.Time `json:"build_staff_update_date" db:"build_staff_update_date"`
}
