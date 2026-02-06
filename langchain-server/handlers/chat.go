package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"langchain-server/agent"
	"langchain-server/mcp"
	"langchain-server/types"

	"github.com/gofiber/fiber/v3"
)

func validateChatRequest(body *types.RequestChatBody) error {
	if body.Credential.Provider == "" {
		return types.NewErrorRequest("Missing provider", 400)
	}

	apiKeyProviders := []string{"openai", "claude", "openrouter"}
	if contains(apiKeyProviders, body.Credential.Provider) && body.Credential.APIKey == nil {
		return types.NewErrorRequest("Missing api key", 401)
	}

	urlProviders := []string{"ollama", "llama_cpp", "vllm"}
	if contains(urlProviders, body.Credential.Provider) && body.Credential.URL == nil {
		return types.NewErrorRequest("Missing url", 401)
	}

	if body.Input == "" {
		return types.NewErrorRequest("Missing body request", 400)
	}

	if len(body.Servers) == 0 {
		return types.NewErrorRequest("No MCP servers provided", 400)
	}

	return nil
}

func ChatHandler(c fiber.Ctx) error {
	var body types.RequestChatBody
	if err := c.Bind().JSON(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validateChatRequest(&body); err != nil {
		if errReq, ok := err.(*types.ErrorRequest); ok {
			return c.Status(errReq.Code).JSON(fiber.Map{
				"error": errReq.Message,
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	availableServers := mcp.CheckServers(body.Servers)
	if len(availableServers) == 0 {
		return c.Status(503).JSON(fiber.Map{
			"error": "No MCP servers available",
		})
	}

	ag, err := agent.CreateLangChainAgent(body.Credential, availableServers, body.SystemPrompt)
	if err != nil {
		if errReq, ok := err.(*types.ErrorRequest); ok {
			return c.Status(errReq.Code).JSON(fiber.Map{
				"error": errReq.Message,
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := c.Context()
	result, err := ag.Invoke(ctx, body.Input)
	if err != nil {
		if errReq, ok := err.(*types.ErrorRequest); ok {
			return c.Status(errReq.Code).JSON(fiber.Map{
				"error": errReq.Message,
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := types.ChatResponse{
		Messages: result.Messages,
		Message:  "",
	}

	if result.Message != nil {
		response.Message = *result.Message
	} else if len(result.Messages) > 0 {
		lastMsg := result.Messages[len(result.Messages)-1]
		response.Message = lastMsg.Content
	}

	return c.JSON(response)
}

func ChatStreamHandler(c fiber.Ctx) error {
	var body types.RequestChatBody
	if err := c.Bind().JSON(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validateChatRequest(&body); err != nil {
		if errReq, ok := err.(*types.ErrorRequest); ok {
			return c.Status(errReq.Code).JSON(fiber.Map{
				"error": errReq.Message,
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	sendEvent := func(data interface{}) error {
		jsonData, _ := json.Marshal(data)
		_, err := fmt.Fprintf(c, "data: %s\n\n", string(jsonData))
		return err
	}

	sendEvent(map[string]interface{}{
		"type":      "start",
		"timestamp": time.Now().Format(time.RFC3339),
		"input":     body.Input,
	})

	availableServers := mcp.CheckServers(body.Servers)
	if len(availableServers) == 0 {
		sendEvent(map[string]interface{}{
			"type":      "error",
			"error":     "No MCP servers available",
			"code":      503,
			"timestamp": time.Now().Format(time.RFC3339),
		})
		return nil
	}

	sendEvent(map[string]interface{}{
		"type":              "servers_checked",
		"available_servers": availableServers,
		"total_servers":     len(body.Servers),
	})

	ag, err := agent.CreateLangChainAgent(body.Credential, availableServers, body.SystemPrompt)
	if err != nil {
		errorCode := 500
		if errReq, ok := err.(*types.ErrorRequest); ok {
			errorCode = errReq.Code
		}
		sendEvent(map[string]interface{}{
			"type":      "error",
			"error":     err.Error(),
			"code":      errorCode,
			"timestamp": time.Now().Format(time.RFC3339),
		})
		return nil
	}

	eventChan := make(chan agent.StreamEvent, 100)
	ctx := c.Context()

	go func() {
		if err := ag.StreamInvoke(ctx, body.Input, eventChan); err != nil {
			errorCode := 500
			if errReq, ok := err.(*types.ErrorRequest); ok {
				errorCode = errReq.Code
			}
			sendEvent(map[string]interface{}{
				"type":      "error",
				"error":     err.Error(),
				"code":      errorCode,
				"timestamp": time.Now().Format(time.RFC3339),
			})
		}
	}()

	for event := range eventChan {
		eventData := map[string]interface{}{
			"type": event.Type,
		}
		if event.Timestamp != "" {
			eventData["timestamp"] = event.Timestamp
		}
		for k, v := range event.Data {
			eventData[k] = v
		}
		if err := sendEvent(eventData); err != nil {
			return err
		}
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
