package checkouts

import (
	"sell/models"
	"sell/repository"

	"github.com/gofiber/fiber/v2"
)

// Checkout - POST
func Checkout(c *fiber.Ctx) error {

	checkoutBody := new(models.CheckoutBody)

	if err := c.BodyParser(checkoutBody); err != nil {
		return c.SendStatus(400)
	}

	checkout, err := repository.MakeCheckout(checkoutBody); 
	
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(checkout)
}