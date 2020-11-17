package models

type Checkout struct {
	Cart Cart `json:cart`
	Order Order `json:"order"`
	UserInfo UserInfo `json:"userInfo"`
	Address Address `json:"address"`
	OrderSell OrderSell `json:"orderSell"`
	ProductDepositMovement []ProductDepositMovement `json:"productDepositMovements"`
}