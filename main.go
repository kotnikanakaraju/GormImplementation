package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an Awesome API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)
	// User endpoints
	app.Post("/api/users", CreateUser)
	app.Get("/api/users", GetUsers)
	app.Get("/api/users/:id",GetUser)
	app.Delete("/api/users/:id",DeleteUser)
	// Product endpoints
	app.Post("/api/products",CreateProduct)
	app.Get("/api/products",GetProducts)
	app.Get("/api/products/:id",GetProduct)
	app.Put("/api/products/:id",UpdateProduct)
	// Order endpoints
	app.Post("/api/orders",CreateOrder)
	app.Get("/api/orders", GetOrders)
	app.Get("/api/orders/:id",GetOrder)
}

func main() {
	ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}