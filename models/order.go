package models

type Order struct {
	ID uint `db:"id,omitempty" json:"id"`
	Type string `db:"type,omitempty" json:"type"`
	Status string `db:"status,omitempty" json:"status"`
	Closed bool `db:"closed,omitempty" json:"closed"`
}