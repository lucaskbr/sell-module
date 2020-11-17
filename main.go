// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	avaibleproducts "sell/controllers/avaible-products"
	"sell/controllers/carts"
	"sell/controllers/checkouts"
	promotedproducts "sell/controllers/promoted-products"
	"sell/database"

	"github.com/gofiber/fiber/v2"
	"github.com/upper/db/v4/adapter/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `corporative_system`,
	Host:     `localhost`,
	User:     `postgres`,
	Password: `docker`,
}


func main() {

	database.Initalize()

	// Fiber instance
	app := fiber.New()

	// Routes

	// PromotedProducts
	app.Get("/promoted-products", promotedproducts.FindAllPromotedProducts)

	// AvaibleProducts
	app.Get("/avaible-products", avaibleproducts.FindAllAvaibleProducts)
	app.Get("/avaible-products/:productDepositId", avaibleproducts.FindProducDepositByID)

	// Carts
	app.Post("/carts", carts.CreateCart)
	app.Get("/carts/:uuid", carts.FindCartByUUID)
	app.Put("/carts/:uuid", carts.AddProductToCart)
	// app.Delete("/carts/:uuid/:productDepositId", carts.AddProductToCart)

	// Checkout
	app.Post("/checkouts", checkouts.Checkout)

	// Start server
	log.Fatal(app.Listen(":3000"))

	database.Close()

}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}