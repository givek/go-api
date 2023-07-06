package main

import (
	"log"

	"github.com/givek/go-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/api/products", handlers.GetAllProducts)

	app.Post("/api/products", handlers.CreateProduct)

	log.Fatal(app.Listen(":8080"))
}
