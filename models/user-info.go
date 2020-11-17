package models

type UserInfo struct {
	ID uint `db:"id,omitempty" json:"id"`
	Email string `db:"email,omitempty" json:"email"`
	Name string `db:"name,omitempty" json:"name"`
	LastName string `db:"last_name,omitempty" json:"lastName"`
	Cpf int `db:"cpf,omitempty" json:"cpf"`
	PhoneNumber string `db:"phone_number,omitempty" json:"phoneNumber"`
}