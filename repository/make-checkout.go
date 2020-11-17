package repository

import (
	"errors"
	"fmt"
	"sell/database"
	"sell/models"
	"time"
)

// MakeCheckout
func MakeCheckout(data *models.CheckoutBody) (interface{}, error) {

  var err error

  cart, err := FindCartByUUID(data.CartUUID)

  if err != nil {
    fmt.Print(err)
    return nil, err
  }

  if (!cart.Active) {
    return nil, errors.New("The cart is already closed")
  }


  order := models.Order{
    Type: "SELL",
    Status: "WAITING FOR PAYMENT",
    Closed: true,
  }

  database.Sess.SQL().
  InsertInto("order").
  Columns(
    "type", 
    "status", 
    "closed",
  ).
  Values(
    order.Type, 
    order.Status, 
    order.Closed,
  ).
  Returning(
    "id",
    "type",
    "status",
    "closed",
  ).
  Iterator().
  Next(&order)

  _, err = CreateOrderHistory(order)

  if (err != nil) {
    fmt.Print(err)
    return nil, err
  }

  userInfo := models.UserInfo{
    Cpf: data.User.Cpf,
    Email: data.User.Email,
    Name: data.User.Name,
    LastName: data.User.Lastname,
    PhoneNumber: data.User.PhoneNumber,
  }

  database.Sess.SQL().
    InsertInto("user_info").
    Columns(     
     "email",
     "name",
     "last_name",
     "cpf",
     "phone_number").
    Values(
     userInfo.Email,
     userInfo.Name,
     userInfo.LastName,
     userInfo.Cpf,
     userInfo.PhoneNumber,
    ).
    Returning(
     "id",
     "email",
     "name",
     "last_name",
     "cpf",
     "phone_number",
    ).
    Iterator().
    Next(&userInfo)

  address := models.Address{
    Cep: data.Address.Cep,
    Number: data.Address.Number,
  }

  database.Sess.SQL().
    InsertInto("address").
    Columns(     
      "cep",
      "number",
    ).
    Values(
      address.Cep, 
      address.Number,
    ).
    Returning(
      "id",
      "cep",
      "number",
    ).
    Iterator().
    Next(&address)

  orderSell := models.OrderSell{
    AddressID: address.ID,
    OrderID: order.ID,
    UserInfoID: userInfo.ID,
  }

  database.Sess.SQL().
    InsertInto("order_sell").
    Columns(     
      "address_id",
      "order_id",
      "user_info_id",
      "cart_id",
    ).
    Values(
      orderSell.AddressID,
      orderSell.OrderID,
      orderSell.UserInfoID,
      cart.ID,
    ).
    Returning(
      "id",
      "address_id",
      "order_id",
      "user_info_id",
      "cart_id",
    ).
    Iterator().
    Next(&orderSell)

  
  var productsDepositMovementList []models.ProductDepositMovement
    
  err = database.Sess.SQL().
    Select(
      "pdm.id AS id",
      "pdm.price AS price",
      "pdm.quantity AS quantity",
      "pdm.product_deposit_id AS product_deposit_id",
    ).
    From("cart c").
    Join("product_deposit_movement pdm").On("pdm.product_movement_id = c.product_movement_id").
    Where("c.uuid", data.CartUUID).And("c.active", true).
    All(&productsDepositMovementList);

  if err != nil {
    fmt.Print(err)
    return nil, err
  }

  res, err := database.Sess.SQL().
  Update("cart").
  Set("active", false).
  Set("updated_at", time.Now()).
  Where("id = ?", cart.ID).
  Exec()

  if err != nil {
    fmt.Print(err)
    return nil, err
  }
  
  fmt.Print(res)
  fmt.Print(productsDepositMovementList)


  fmt.Print(order)
  fmt.Print(userInfo)
  fmt.Print(orderSell)


  checkout := models.Checkout{
    Cart: cart,
    Order: order,
    OrderSell: orderSell,
    ProductDepositMovement: productsDepositMovementList,
    UserInfo: userInfo,
    Address: address,
  }

  return checkout, err

}