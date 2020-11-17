package repository

import (
	"sell/database"
	"sell/models"
	"time"
)

// CreateProductMovement
func CreatePickupReservationProductMovement() (models.ProductMovement, error) {

	var document models.Document

	database.Sess.SQL().
  	InsertInto("document").
  	Columns(
  	  "path", 
  	  "type",
  	).
  	Values(
  	  "PICKUP-RESERVATION",
  	  "PICKUP-RESERVATION",
  	).
  	Returning(
  	  "id",
  	).
  	Iterator().
  	Next(&document)

	var productMovement models.ProductMovement

	database.Sess.SQL().
  	InsertInto("product_movement").
  	Columns(
  	  "type", 
			"created_at",
			"document_id",
  	).
  	Values(
  	  "PICKUP-RESERVATION",
			time.Now(),
			document.ID,
  	).
  	Returning(
			"id",
			"type",
			"created_at",
			"updated_at",
			"document_id",
  	).
  	Iterator().
		Next(&productMovement)
		
  return productMovement, nil
}