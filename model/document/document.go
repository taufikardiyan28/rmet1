package DocumentModel

import "time"

type Document struct {
	DocID           int64     `db:"doc_id" json:"doc_id"`
	DocBuildID      int64     `db:"doc_build_id" json:"doc_build_id"`
	DocCreateBy     string    `db:"doc_createe_by" json:"doc_create_by"`
	DocCreateDate   time.Time `db:"doc_created_date" json:"doc_createe_date"`
	DocDelStatus    string    `db:"doc_del_status" json:"doc_del_status"`
	DocFileName     string    `db:"doc_file_name" json:"doc_file_name"`
	DocOriginalName string    `db:"doc_original_name" json:"doc_original_name"`
	DocPath         string    `db:"doc_path" json:"doc_path"`
	DocUpdateBy     string    `db:"doc_update_by" json:"doc_update_by"`
	DocUpdateDate   time.Time `db:"doc_update_date" json:"doc_update_date"`
}
