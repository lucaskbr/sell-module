package repository

import (
	"fmt"
	"sell/database"
	"sell/models"

	"github.com/google/uuid"
)

// CreateCart
func CreateCart() (string, error) {


  productMovement, err := CreatePickupReservationProductMovement()

  if err != nil {
    fmt.Print(err)
    return "", err
  }
	
  uuid, err := uuid.NewUUID()

  col := database.Sess.Collection("cart")

  cart := models.Cart{
    UUID:     uuid.String(),
    ProductMovementID: productMovement.ID,
  }
  
  _, err = col.Insert(cart)

  return uuid.String(), err
}