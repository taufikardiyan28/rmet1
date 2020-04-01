package UserModel

import "time"

type User struct {
	UserID               int64     `json:"user_id" db:"user_id"`
	UserCabang           string    `json:"user_cabang" db:"user_cabang"`
	UserCreateBy         string    `json:"user_create_by" db:"user_create_by"`
	UserCreateDate       time.Time `json:"user_create_date" db:"user_create_date"`
	UserDelStatus        string    `json:"user_del_status" db:"user_del_status"`
	UserDeptID           int64     `json:"user_dept_id" db:"user_dept_id"`
	UserEmail            string    `json:"user_email" db:"user_email"`
	UserFirstname        string    `json:"user_firstname" db:"user_firstname"`
	UserGender           string    `json:"user_gender" db:"user_gender"`
	UserHashkey          string    `json:"user_hashkey" db:"user_hashkey"`
	UserJoindate         time.Time `json:"user_joindate" db:"user_joindate"`
	UserLastname         string    `json:"user_lastname" db:"user_lastname"`
	UserLoginHashkey     string    `json:"user_login_hashkey" db:"user_login_hashkey"`
	UserLoginHashpass    string    `json:"user_login_hashpass" db:"user_login_hashpass"`
	UserLoginID          string    `json:"user_login_id" db:"user_login_id"`
	UserLoginMultidevice int64     `json:"user_login_multidevice" db:"user_login_multidevice"`
	UserName             string    `json:"user_name" db:"user_name"`
	UserNik              string    `json:"user_nik" db:"user_nik"`
	UserPassword         string    `json:"user_password" db:"user_password"`
	UserPhone            string    `json:"user_phone" db:"user_phone"`
	UserPhoto            string    `json:"user_photo" db:"user_photo"`
	UserRoleID           int64     `json:"user_role_id" db:"user_role_id"`
	UserUpdateBy         string    `json:"user_update_by" db:"user_update_by"`
	UserUpdateDate       time.Time `json:"user_update_date" db:"user_update_date"`
}
