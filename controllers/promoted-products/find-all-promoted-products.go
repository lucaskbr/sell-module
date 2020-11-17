package promotedproducts

import (
	"sell/repository"

	"github.com/gofiber/fiber/v2"
)

// FindAllPromotedProducts - GET
func FindAllPromotedProducts(c *fiber.Ctx) error {

	promotedProducts, err :=	repository.FindAllPromotedProducts()

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(promotedProducts)
}