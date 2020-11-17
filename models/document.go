package models

type Document struct {
	ID uint `db:"id,omitempty" json:"id"`
	Path string `db:"path,omitempty" json:"path"`
	Type string `db:"type,omitempty" json:"type"`
}
