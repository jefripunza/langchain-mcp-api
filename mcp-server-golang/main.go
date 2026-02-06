package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
