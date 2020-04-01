package DocumentModel

import (
	db "github.com/taufikardiyan28/rmet1/db"
)

type Document struct {
	DocID           int64         `db:"doc_id" json:"doc_id"`
	DocBuildID      db.NullInt64  `db:"doc_build_id" json:"doc_build_id"`
	DocCreateBy     db.NullString `db:"doc_create_by" json:"doc_create_by"`
	DocCreateDate   db.NullTime   `db:"doc_create_date" json:"doc_create_date"`
	DocDelStatus    db.NullString `db:"doc_del_status" json:"doc_del_status"`
	DocFileName     db.NullString `db:"doc_file_name" json:"doc_file_name"`
	DocOriginalName db.NullString `db:"doc_original_name" json:"doc_original_name"`
	DocPath         db.NullString `db:"doc_path" json:"doc_path"`
	DocUpdateBy     db.NullString `db:"doc_update_by" json:"doc_update_by"`
	DocUpdateDate   db.NullTime   `db:"doc_update_date" json:"doc_update_date"`
}
