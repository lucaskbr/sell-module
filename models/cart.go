package models

type Cart struct {
	ID uint `db:"id,omitempty" json:"id"`
	UUID string `db:"uuid,omitempty" json:"uuid"`
	Active bool `db:"active,omitempty" json:"active"`
//	ProductDepositCart []ProductDepositCart `json:"products"`
	ProductMovementID uint `db:"product_movement_id,omitempty" json:"productMovementId"`
	ProductDepositMovements []ProductDepositMovement `json:"products"`
}