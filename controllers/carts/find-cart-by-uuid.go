package carts

import (
	"sell/repository"

	"github.com/gofiber/fiber/v2"
)

// FindCartById - GET
func FindCartByUUID(c *fiber.Ctx) error {

	uuid := c.Params("uuid")

	cart, err := repository.FindCartByUUID(uuid)

	if err != nil {
		return c.SendStatus(404)
	}

	return c.JSON(cart)
}