package repository

import (
	"errors"
	"fmt"
	"sell/database"
	"time"
)

// AddProductToCart
func AddProductToCart(cartUUID string, productDepositID uint, quantity int, action string) (interface{}, error) {


  productDeposit, err := FindProductDepositByID(productDepositID)

  if err != nil {
    fmt.Println(err)
    return nil, err
  }

  productDepositMovementCart, _ := FindProductInCart(cartUUID, productDepositID)

  var quantityToUpdateMovement int
  var quantityToUpdateDeposit int


  if productDepositMovementCart.ProductDepositID != 0 {

    if (action == "ADD") {


      if quantity > productDeposit.Quantity {
        return nil, errors.New("ProductDeposit exceeds the quantity to a pickup reservation")
      }

      quantityToUpdateDeposit = productDeposit.Quantity  - quantity
      quantityToUpdateMovement = productDepositMovementCart.Quantity + quantity

      fmt.Println(quantityToUpdateDeposit)
      fmt.Println(quantityToUpdateMovement)
      fmt.Println(productDepositMovementCart.Quantity)

    } else {

      if productDepositMovementCart.Quantity < quantity {
        return nil, errors.New("Cannot remove what you dont have from cart")
      }

      quantityToUpdateDeposit = productDeposit.Quantity + quantity

      quantityToUpdateMovement = productDepositMovementCart.Quantity - quantity

      if (quantityToUpdateMovement < 0) {
        quantityToUpdateMovement = 0
      }
    }

    
    res, err := database.Sess.SQL().
      Update("product_deposit_movement").
      Set("quantity", quantityToUpdateMovement).
      Set("updated_at", time.Now()).
      Where("id = ?", productDepositMovementCart.ID).
      Exec()

    if err != nil {
      fmt.Println(err)
      return nil, err
    }

    res, err = database.Sess.SQL().
      Update("product_deposit").
      Set("quantity", quantityToUpdateDeposit).
      Set("updated_at", time.Now()).
      Where("id = ?", productDepositID).
      Exec()

    if err != nil {
      fmt.Println(err)
      return nil, err
    }

    return res, nil

  }

  cart, err := FindCartByUUID(cartUUID)

  if (productDeposit.Quantity < quantity) {
    return nil, err
  }

  quantityToUpdateDeposit = productDeposit.Quantity - quantity


  res, err := database.Sess.SQL().
    Update("product_deposit").
    Set("quantity", quantityToUpdateDeposit).
    Set("updated_at", time.Now()).
    Where("id = ?", productDepositID).
    Exec()

  res, err = database.Sess.SQL().
    InsertInto("product_deposit_movement").
    Columns(
      "product_deposit_id",
      "price",
      "quantity",
      "product_movement_id",
    ).
    Values(
      productDepositID,
      productDeposit.Price,
      quantity,
      cart.ProductMovementID,
    ).
    Exec()

 return res, nil
}