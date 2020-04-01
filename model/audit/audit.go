package AuditModel

import (
	db "github.com/taufikardiyan28/rmet1/db"
)

type Audit struct {
	AuditID          int64         `json:"audit_id" db:"audit_id"`
	AuditAuditorName db.NullString `json:"audit_auditor_name" db:"audit_auditor_name"`
	AuditBuildID     db.NullInt64  `json:"audit_build_id" db:"audit_build_id"`
	AuditBuildName   db.NullString `json:"audit_build_name" db:"audit_build_name"`
	AuditCreateBy    db.NullString `json:"audit_create_by" db:"audit_create_by"`
	AuditCreateDate  db.NullTime   `json:"audit_create_date" db:"audit_create_date"`
	AuditDate        db.NullTime   `json:"audit_date" db:"audit_date"`
	AuditDelStatus   db.NullString `json:"audit_del_status" db:"audit_del_status"`
	AuditUpdateBy    db.NullString `json:"audit_update_by" db:"audit_update_by"`
	AuditUpdateDate  db.NullTime   `json:"audit_update_date" db:"audit_update_date"`
}
