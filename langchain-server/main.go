package main

import (
	"log"

	"langchain-server/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Langchain Server",
		TrustProxy:    true,
		// ReduceMemoryUsage: true,
	})

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} - ${method} ${path}\n",
	}))
	app.Use(requestid.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "🤖 LangChain Server is running",
			"version": "1.0.0",
		})
	})

	app.Post("/chat", handlers.ChatHandler)
	app.Post("/chat/stream", handlers.ChatStreamHandler)

	log.Println("🤖 LangChain Server running on http://localhost:6000")
	log.Fatal(app.Listen(":6000"))
}
