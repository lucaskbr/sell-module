package models

// Product
type product struct {
	ID int `db:"product_id" json:"id"`
	ProductName string `db:"product_name" json:"productName"`
}

// Provider
type provider struct {
	ID int `db:"provider_id" json:"id"`
	Name string `db:"provider_name" json:"name"`
}

// ProductDeposit
type ProductDeposit struct {
	ProductDepositID int `db:"id" json:"productDepositId"`
	Price float64 `db:"price" json:"price"`
	Quantity int `db:"quantity" json:"quantity"`
	Product product `json:"product"`
	Provider provider `json:"provider"`
}