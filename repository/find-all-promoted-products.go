package repository

import (
	"sell/database"
	"time"
)

// Product
type ProductDPP struct {
	ID int `db:"product_id" json:"id"`
	ProductName string `db:"product_name" json:"productName"`
	Price float64 `db:"price" json:"price"`
	Quantity float64 `db:"quantity" json:"quantity"`
}

// Provider
type ProviderDPP struct {
	ID int `db:"provider_id" json:"id"`
	Name string `db:"provider_name" json:"name"`
}

// ProductDepositPromote
type ProductDepositPromote struct {
	ID int `db:"id" json:"id"`
	ProductDepositPromoteId int `db:"product_deposit_promote_id" json:"productDepositPromoteId"`
	Until time.Time `db:"until" json:"until"`
	Product ProductDPP `json:"product"`
	Provider ProviderDPP `json:"provider"`
}

// FindPromotedProducts
func FindAllPromotedProducts() ([]ProductDepositPromote, error) {

	q := database.Sess.
	SQL().
	Select(
		// product_deposit_promote
		"pdp.id AS product_deposit_promote_id",
		"pdp.until AS until",
		// product_deposit
		"pd.deposit_id AS id",
		"pd.price AS price",
		"pd.quantity AS quantity",

		// Product
		"product.id AS product_id",
		"product.name AS product_name",

		// Provider
		"provider.id AS provider_id",
		"provider.name AS provider_name",
		).
	From("product_deposit_promote AS pdp").
	Join("product_deposit AS pd").On("pdp.product_deposit_id = pd.id").
	Join("product_provider AS pp").On("pd.product_provider_id = pp.id").
	Join("product").On("pp.product_id = product.id").
	Join("provider").On("pp.provider_id = provider.id").
	Where("pdp.until >", time.Now())

	var productsDepositPromote []ProductDepositPromote
	
	err := q.All(&productsDepositPromote);
	
	return productsDepositPromote, err

}