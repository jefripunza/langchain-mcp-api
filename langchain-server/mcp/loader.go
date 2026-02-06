package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"langchain-mcp-api/types"
	"langchain-mcp-api/utils"

	"github.com/tmc/langchaingo/tools"
)

type MCPTool struct {
	name        string
	description string
	mcpURL      string
	toolDef     types.Tool
}

func (t *MCPTool) Name() string {
	return t.name
}

func (t *MCPTool) Description() string {
	return t.description
}

func (t *MCPTool) Call(ctx context.Context, input string) (string, error) {
	var args map[string]interface{}
	if err := json.Unmarshal([]byte(input), &args); err != nil {
		return "", fmt.Errorf("failed to parse input: %w", err)
	}

	result, err := InvokeTool(t.mcpURL, t.name, args)
	if err != nil {
		return "", err
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(resultJSON), nil
}

func LoadMCPToolsAsLangChain(requestID string, mcpServers []string) ([]tools.Tool, []types.Tool, error) {
	utils.VerbosePrintf("\n[%s]🔌 [MCP] Loading tools from MCP servers...\n", requestID)
	var langchainTools []tools.Tool
	var toolDefs []types.Tool

	for idx, mcpURL := range mcpServers {
		utils.VerbosePrintf("[%s]   [%d/%d] Fetching from: %s\n", requestID, idx+1, len(mcpServers), mcpURL)
		serverTools, err := fetchToolsFromServer(mcpURL)
		if err != nil {
			utils.VerbosePrintf("      ❌ Failed: %v\n", err)
			continue
		}
		utils.VerbosePrintf("[%s]      ✅ Loaded %d tools\n", requestID, len(serverTools))

		for _, toolDef := range serverTools {
			utils.VerbosePrintf("[%s]         - %s: %s\n", requestID, toolDef.Name, toolDef.Description)
			mcpTool := &MCPTool{
				name:        toolDef.Name,
				description: toolDef.Description,
				mcpURL:      mcpURL,
				toolDef:     toolDef,
			}
			langchainTools = append(langchainTools, mcpTool)
			toolDefs = append(toolDefs, toolDef)
		}
	}

	utils.VerbosePrintf("\n[%s]✅ [MCP] Total tools loaded: %d\n", requestID, len(langchainTools))
	return langchainTools, toolDefs, nil
}

func fetchToolsFromServer(mcpURL string) ([]types.Tool, error) {
	resp, err := http.Get(mcpURL + "/mcp/tools")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch tools: %s", resp.Status)
	}

	var tools []types.Tool
	if err := json.NewDecoder(resp.Body).Decode(&tools); err != nil {
		return nil, err
	}

	return tools, nil
}

func InvokeTool(mcpURL string, toolName string, args map[string]interface{}) (interface{}, error) {
	reqBody := types.ToolInvokeRequest{
		Name:      toolName,
		Arguments: args,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(mcpURL+"/mcp/invoke", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("tool invocation failed: %s - %s", resp.Status, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func CheckServers(requestID string, mcpServers []string) []string {
	utils.VerbosePrintf("\n[%s]🏥 [MCP] Checking server health...\n", requestID)
	var availableServers []string

	for idx, serverURL := range mcpServers {
		utils.VerbosePrintf("[%s]   [%d/%d] Checking: %s\n", requestID, idx+1, len(mcpServers), serverURL)
		resp, err := http.Get(serverURL + "/health")
		if err != nil {
			utils.VerbosePrintf("[%s]      ❌ Not available: %v\n", requestID, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			availableServers = append(availableServers, serverURL)
			utils.VerbosePrintf("[%s]      ✅ Healthy (status %d)\n", requestID, resp.StatusCode)
		} else {
			utils.VerbosePrintf("[%s]      ⚠️  Unhealthy (status %d)\n", requestID, resp.StatusCode)
		}
	}

	utils.VerbosePrintf("\n[%s]✅ [MCP] Available servers: %d/%d\n", requestID, len(availableServers), len(mcpServers))
	return availableServers
}
