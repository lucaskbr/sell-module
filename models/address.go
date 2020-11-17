package models

type Address struct {
	ID uint `db:"id,omitempty" json:"id"`
	Cep string `db:"cep,omitempty" json:"cep"`
	Number int `db:"number,omitempty" json:"number"`
}