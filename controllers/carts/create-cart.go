package carts

import (
	"sell/repository"

	"github.com/gofiber/fiber/v2"
)

// FindCartById - GET
func CreateCart(c *fiber.Ctx) error {

	uuid, err := repository.CreateCart()

	if err != nil {
		return c.SendStatus(500)
	}

	res := struct { UUID string `json:"uuid"` } { uuid }

	return c.JSON(res)
}