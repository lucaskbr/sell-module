package avaibleproducts

import (
	"sell/repository"

	"github.com/gofiber/fiber/v2"
)

// FindAllAvaibleProducts - GET
func FindAllAvaibleProducts(c *fiber.Ctx) error {

	avaibleProducts, err := repository.FindAllAvaibleProducts()

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(avaibleProducts)
}