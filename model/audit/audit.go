package AuditModel

import "time"

type Audit struct {
	AuditID          int64     `json:"audit_id" db:"audit_id"`
	AuditAuditorName string    `json:"audit_auditor_name" db:"audit_auditor_name"`
	AuditBuildID     int64     `json:"audit_build_id" db:"audit_build_id"`
	AuditBuildName   string    `json:"audit_build_name" db:"audit_build_name"`
	AuditCreateBy    string    `json:"audit_create_by" db:"audit_create_by"`
	AuditCreateDate  time.Time `json:"audit_create_date" db:"audit_create_date"`
	AuditDate        time.Time `json:"audit_date" db:"audit_date"`
	AuditDelStatus   string    `json:"audit_del_status" db:"audit_del_status"`
	AuditUpdateBy    string    `json:"audit_update_by" db:"audit_update_by"`
	AuditUpdateDate  time.Time `json:"audit_update_date" db:"audit_update_date"`
}
