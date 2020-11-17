package repository

import (
	"sell/database"
	"sell/models"
)

// FindAllAvaibleProducts
func FindAllAvaibleProducts() ([]models.ProductDeposit, error) {

	q := database.Sess.
	SQL().
	Select(
		// product_deposit
		"pd.id AS id",
		"pd.deposit_id AS deposit_id",
		"pd.price AS price",
		"pd.quantity AS quantity",

		// Product
		"product.id AS product_id",
		"product.name AS product_name",

		// Provider
		"provider.id AS provider_id",
		"provider.name AS provider_name",
		).
	From("product_deposit AS pd").
	Join("product_provider AS pp").On("pd.product_provider_id = pp.id").
	Join("product").On("pp.product_id = product.id").
	Join("provider").On("pp.provider_id = provider.id").
	Where("pd.quantity >", 0)

	var productsDeposit []models.ProductDeposit
	
	err := q.All(&productsDeposit);

	return productsDeposit, err

}