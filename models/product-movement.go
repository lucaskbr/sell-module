package models

import "time"

// ProductMovement
type ProductMovement struct {
	ID uint `db:"id" json:"id"`
	Type string `db:"type" json:"type"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	DocumentID uint `db:"document_id" json:"documentId"`
}