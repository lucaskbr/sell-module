package carts

import (
	"sell/repository"

	"github.com/gofiber/fiber/v2"
)

type AddProductToCartBody struct {
	ProductDepositID uint `json:"productDepositId"`
	Quantity int `json:"quantity"`
	Action string `json:"action"`  // ADD - REMOVE
}

// AddProductToCart - PUT
func AddProductToCart(c *fiber.Ctx) error {

	var UUID = c.Params("uuid")

	body := new(AddProductToCartBody)
	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(400)
	}

	if body.Quantity < 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "Quantity needs to be 0 or above")
	}


	_, err := repository.FindCartByUUID(UUID)

	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "The cart searched doesnt exists or is closed")
	}


	_, err = repository.AddProductToCart(UUID, body.ProductDepositID, body.Quantity, body.Action); 
	
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	return c.SendStatus(204)
}