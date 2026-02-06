package main

import (
	"log"
	"mcp-server/registry"
	"mcp-server/types"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Get("/mcp/tools", func(c fiber.Ctx) error {
		toolsResponse := make([]types.ToolResponse, len(registry.Tools))
		for i, tool := range registry.Tools {
			toolsResponse[i] = types.ToolResponse{
				Name:        tool.Name,
				Description: tool.Description,
				Parameters:  tool.Parameters,
			}
		}
		return c.JSON(toolsResponse)
	})

	app.Post("/mcp/invoke", func(c fiber.Ctx) error {
		var body struct {
			Name      string                 `json:"name"`
			Arguments map[string]interface{} `json:"arguments"`
		}

		if err := c.Bind().Body(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		tool := registry.FindTool(body.Name)
		if tool == nil {
			return c.Status(404).JSON(fiber.Map{"error": "Tool not found"})
		}

		if tool.Handler == nil {
			return c.Status(400).JSON(fiber.Map{"error": "Tool handler not found"})
		}

		result, err := tool.Handler(body.Arguments)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(result)
	})

	log.Println("🧠 MCP Server running on http://localhost:4040")
	log.Fatal(app.Listen(":4040"))
}
