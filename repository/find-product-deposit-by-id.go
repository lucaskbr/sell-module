package repository

import (
	"sell/database"
	"sell/models"
)

// FindProductDepositByID
func FindProductDepositByID(id uint) (models.ProductDeposit, error) {

	var productDeposit models.ProductDeposit

	err := database.Sess.
	SQL().
	Select(
		// product_deposit
		"pd.id AS id",
		"pd.average_price AS average_price",
		"pd.price AS price",
		"pd.quantity AS quantity",
		"pd.deposit_id AS deposit_id",
		"pd.product_provider_id AS product_provider_id",
		).
	From("product_deposit AS pd").
	Where("pd.id =", id).One(&productDeposit)

	return productDeposit, err

}