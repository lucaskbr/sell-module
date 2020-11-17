package models

type OrderSell struct {
	ID uint `db:"id,omitempty" json:"id"`
	OrderID uint `db:"order_id,omitempty" json:"orderId"`
	AddressID uint `db:"address_id,omitempty" json:"addressId"`
	UserInfoID uint `db:"user_info_id,omitempty" json:"userInfoId"`
}