package avaibleproducts

import (
	"sell/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// FindProducDepositByID - GET
func FindProducDepositByID(c *fiber.Ctx) error {

	var ID64, err = strconv.ParseUint(c.Params("productDepositId"), 10, 32)

	if err != nil {
		return c.SendStatus(500)
	}

	ID := uint(ID64)

	productDeposit, err := repository.FindProductDepositByID(ID)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(productDeposit)
}