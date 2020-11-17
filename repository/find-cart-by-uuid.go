package repository

import (
	"fmt"
	"sell/database"
	"sell/models"
)

// FindCartByUUID
func FindCartByUUID(uuid string) (models.Cart, error) {

	var cart models.Cart

	err := database.Sess.SQL().
	Select(
		"c.id AS id",
		"c.uuid AS uuid",
		"c.active AS active",
		"c.product_movement_id AS product_movement_id",
		).
	From("cart AS c").
	Where("uuid =", uuid).
	One(&cart)

	if err != nil {
    return cart, err
	}
	

	var productDepositMovements []models.ProductDepositMovement

	err = database.Sess.SQL().
		Select(
			"pdm.id AS id",
			"pdm.price AS price",
			"pdm.quantity AS quantity",
			"pdm.product_deposit_id AS product_deposit_id",
			"pdm.product_movement_id AS product_movement_id",
		).
		From("product_deposit_movement AS pdm").
		Where("pdm.product_movement_id =", cart.ProductMovementID).
		All(&productDepositMovements)

	if err != nil {
		fmt.Print(err)
    return cart, err
	}

	cart.ProductDepositMovements = productDepositMovements

	return cart, nil


}