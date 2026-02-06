package main

import (
	"log"
	"time"

	"langchain-mcp-api/handlers"

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
		AppName:       "Langchain MCP API",
		TrustProxy:    true,
		// ReduceMemoryUsage: true,
	})

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${requestID}] ${date} ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
		CustomTags: map[string]logger.LogFunc{
			"requestID": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				reqID := requestid.FromContext(c)
				return output.WriteString(reqID)
			},
			"date": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(time.Now().Format("2006-01-02"))
			},
		},
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "🤖 LangChain MCP API is running",
			"version": "1.0.0",
		})
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/chat", handlers.ChatHandler)
	app.Post("/chat/stream", handlers.ChatStreamHandler)

	log.Fatal(app.Listen("0.0.0.0:6000"))
}
