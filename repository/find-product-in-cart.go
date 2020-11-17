package repository

import (
	"fmt"
	"sell/database"
)

type FindProductInCartResult struct {
	ID uint `db:"id,omitempty" json:"id"`
	ProductDepositID uint `db:"product_deposit_id,omitempty" json:"productDepositId"`
	CartID uint `db:"cart_id,omitempty" json:"CartId"`
	Quantity int `db:"quantity,omitempty" json:"quantity"`
}

// FindProductInCart
func FindProductInCart(uuid string, productDepositID uint) (FindProductInCartResult, error) {

	q := database.Sess.SQL().
	Select(
		"c.id AS cart_id",
		"pdm.id AS id",
		"pdm.product_deposit_id AS product_deposit_id",
		"pdm.quantity AS quantity",
		).
	From("cart AS c").
	Join("product_movement AS pm").On("c.product_movement_id = pm.id").
	Join("product_deposit_movement AS pdm").On("pm.id = pdm.product_movement_id").
	Where("c.uuid = ? AND pdm.product_deposit_id = ?", uuid, productDepositID)


	var findProductInCartResult FindProductInCartResult
	err := q.One(&findProductInCartResult)

	if err != nil {
		fmt.Print(findProductInCartResult)
    return findProductInCartResult, err
	}

	return findProductInCartResult, nil

}