package repository

import (
	"sell/database"
	"sell/models"
)

// CreateOrderHistory
func CreateOrderHistory(order models.Order) (models.OrderHistory, error) {
	
  col := database.Sess.Collection("order_history")

  orderHistory := models.OrderHistory{
		Status: order.Status,
		Type: order.Type,
		Closed: order.Closed,
		OrderID: order.ID,
  }
  
	_, err := col.Insert(orderHistory)
	
  return orderHistory, err
}