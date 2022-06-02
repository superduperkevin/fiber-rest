package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/superduperkevin/fiber-rest/database"
	"github.com/superduperkevin/fiber-rest/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my api")
}

func setupRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api/", welcome)

	//user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	//product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	//order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
	app.Delete("/api/orders/:id", routes.DeleteOrder)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}