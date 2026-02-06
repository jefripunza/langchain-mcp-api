package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"langchain-mcp-api/llm"
	"langchain-mcp-api/mcp"
	"langchain-mcp-api/types"
	"langchain-mcp-api/utils"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/tools"
)

type LangChainAgent struct {
	executor      *agents.Executor
	llmClient     *llm.LangChainClient
	tools         []tools.Tool
	toolDefs      []types.Tool
	mcpServers    []string
	systemPrompt  *string
	supportsTools bool
	provider      string
}

func CreateLangChainAgent(
	requestID string,
	credential types.RequestChatCredential,
	mcpServers []string,
	systemPrompt *string,
) (*LangChainAgent, error) {
	utils.VerbosePrintf("\n[%s]📦 [AGENT] Creating LangChain Agent...\n", requestID)
	utils.VerbosePrintf("[%s]   Provider: %s\n", requestID, credential.Provider)
	if credential.Model != nil {
		utils.VerbosePrintf("[%s]   Model: %s\n", requestID, *credential.Model)
	}
	utils.VerbosePrintf("[%s]   MCP Servers: %d\n", requestID, len(mcpServers))

	langchainTools, toolDefs, err := mcp.LoadMCPToolsAsLangChain(requestID, mcpServers)
	if err != nil {
		return nil, err
	}
	utils.VerbosePrintf("[%s]   ✅ Loaded %d tools from MCP servers\n", requestID, len(langchainTools))

	llmClient, err := llm.CreateLangChainLLM(requestID, credential)
	if err != nil {
		return nil, err
	}

	utils.VerbosePrintf("\n[%s]🤖 Using LLM provider: %s", requestID, credential.Provider)
	if llmClient.SupportsTools {
		utils.VerbosePrintf(" (with native tool calling)\n")
	} else {
		utils.VerbosePrintf(" (with manual tool calling)\n")
	}

	agent := &LangChainAgent{
		llmClient:     llmClient,
		tools:         langchainTools,
		toolDefs:      toolDefs,
		mcpServers:    mcpServers,
		systemPrompt:  systemPrompt,
		supportsTools: llmClient.SupportsTools,
		provider:      credential.Provider,
	}

	if llmClient.SupportsTools {
		utils.VerbosePrintf("[%s]   🔧 Initializing agent executor with native tool calling...\n", requestID)
		executor, err := agents.Initialize(
			llmClient.LLM,
			langchainTools,
			agents.ZeroShotReactDescription,
			agents.WithMemory(memory.NewConversationBuffer()),
		)
		if err != nil {
			return nil, err
		}
		agent.executor = executor
		utils.VerbosePrintf("[%s]   ✅ Agent executor initialized\n", requestID)
	} else {
		utils.VerbosePrintf("[%s]   🔧 Using manual tool calling mode\n", requestID)
	}

	utils.VerbosePrintf("[%s]✅ [AGENT] Agent created successfully\n", requestID)
	return agent, nil
}

func (a *LangChainAgent) Invoke(requestID string, ctx context.Context, input string) (*types.AgentState, error) {
	utils.VerbosePrintf("\n[%s]🚀 [INVOKE] Starting agent invocation...\n", requestID)
	utils.VerbosePrintf("[%s]   Input: %s\n", requestID, input)

	state := &types.AgentState{
		Input:    input,
		Messages: []types.Message{},
	}

	if a.supportsTools && a.executor != nil {
		utils.VerbosePrintf("[%s]   🔄 Using native tool calling executor...\n", requestID)
		result, err := chains.Run(ctx, a.executor, input)
		if err != nil {
			utils.VerbosePrintf("[%s]   ❌ Error: %v\n", requestID, err)
			return nil, err
		}

		utils.VerbosePrintf("[%s]   ✅ Response: %s\n", requestID, result)
		state.Message = &result
		state.Messages = append(state.Messages, types.Message{
			Role:    "assistant",
			Content: result,
		})

		return state, nil
	}

	utils.VerbosePrintf("[%s]   🔄 Using manual tool calling mode...\n", requestID)
	maxIterations := 10
	iteration := 0

	for iteration < maxIterations {
		iteration++
		utils.VerbosePrintf("\n[%s]   🔁 [ITERATION %d/%d]\n", requestID, iteration, maxIterations)

		messages := a.buildMessages(requestID, state)
		utils.VerbosePrintf("[%s]      📝 Built %d messages for LLM\n", requestID, len(messages))
		utils.VerbosePrintf("[%s]      🤖 Calling LLM...\n", requestID)

		content, err := a.llmClient.GenerateContent(requestID, ctx, messages)
		if err != nil {
			utils.VerbosePrintf("[%s]      ❌ LLM Error: %v\n", requestID, err)
			return nil, err
		}
		utils.VerbosePrintf("[%s]      ✅ LLM Response (%d chars)\n", requestID, len(content))

		// Parse tool calls FIRST before truncation
		response := &types.Message{
			Role:    "assistant",
			Content: content,
		}
		response = a.parseManualToolCalls(response)
		state.Messages = append(state.Messages, *response)

		if len(response.ToolCalls) == 0 {
			utils.VerbosePrintf("[%s]      ✅ No tool calls detected - final response\n", requestID)
			lastMessage := response.Content
			state.Message = &lastMessage
			break
		}

		utils.VerbosePrintf("[%s]      🔧 Detected %d tool call(s)\n", requestID, len(response.ToolCalls))
		for i, tc := range response.ToolCalls {
			utils.VerbosePrintf("[%s]         %d. %s(%v)\n", requestID, i+1, tc.Name, tc.Args)
		}

		utils.VerbosePrintf("[%s]      ⚙️  Executing tools...\n", requestID)
		toolMessages, err := a.executeTools(requestID, ctx, response.ToolCalls)
		if err != nil {
			utils.VerbosePrintf("[%s]      ❌ Tool execution error: %v\n", requestID, err)
			return nil, err
		}
		utils.VerbosePrintf("[%s]      ✅ Tools executed successfully (%d results)\n", requestID, len(toolMessages))

		state.Messages = append(state.Messages, toolMessages...)
	}

	utils.VerbosePrintf("[%s]✅ [INVOKE] Agent invocation completed\n", requestID)
	return state, nil
}

func (a *LangChainAgent) StreamInvoke(requestID string, ctx context.Context, input string, eventChan chan<- StreamEvent) error {
	defer close(eventChan)

	state := &types.AgentState{
		Input:    input,
		Messages: []types.Message{},
	}

	eventChan <- StreamEvent{
		Type:      "start",
		Timestamp: time.Now().Format(time.RFC3339),
		Data: map[string]interface{}{
			"input": input,
		},
	}

	if a.supportsTools && a.executor != nil {
		result, err := chains.Run(ctx, a.executor, input)
		if err != nil {
			return err
		}

		eventChan <- StreamEvent{
			Type: "message_chunk",
			Data: map[string]interface{}{
				"chunk":     result,
				"is_final":  true,
				"timestamp": time.Now().Format(time.RFC3339),
			},
		}

		eventChan <- StreamEvent{
			Type: "done",
			Data: map[string]interface{}{
				"done":      true,
				"timestamp": time.Now().Format(time.RFC3339),
			},
		}

		return nil
	}

	maxIterations := 10
	iteration := 0
	stepCount := 0

	for iteration < maxIterations {
		iteration++

		eventChan <- StreamEvent{
			Type: "node_execution",
			Data: map[string]interface{}{
				"node":      "agent_start",
				"step":      stepCount,
				"timestamp": time.Now().Format(time.RFC3339),
			},
		}
		stepCount++

		messages := a.buildMessages(requestID, state)
		contentChan, errChan := a.llmClient.StreamGenerateContent(requestID, ctx, messages)

		accumulatedContent := ""
		isInThinkingMode := false
		isInMessageMode := false
		hasStartedStreaming := false
		thinkingBuffer := ""
		messageBuffer := ""

		for {
			select {
			case chunk, ok := <-contentChan:
				if !ok {
					goto StreamDone
				}

				accumulatedContent += chunk

				if a.provider == "llama_cpp" {
					if strings.Contains(accumulatedContent, "<thinking>") && !isInThinkingMode {
						isInThinkingMode = true
						eventChan <- StreamEvent{
							Type:      "thinking_start",
							Timestamp: time.Now().Format(time.RFC3339),
						}
					}

					if isInThinkingMode && !strings.Contains(accumulatedContent, "</thinking>") {
						thinkingChunk := strings.ReplaceAll(chunk, "<thinking>", "")
						if thinkingChunk != "" {
							thinkingBuffer += thinkingChunk
							eventChan <- StreamEvent{
								Type: "thinking_chunk",
								Data: map[string]interface{}{
									"chunk":     thinkingChunk,
									"is_final":  false,
									"timestamp": time.Now().Format(time.RFC3339),
								},
							}
						}
						continue
					}

					if strings.Contains(accumulatedContent, "</thinking>") && isInThinkingMode {
						re := regexp.MustCompile(`<thinking>([\s\S]*?)</thinking>`)
						matches := re.FindStringSubmatch(accumulatedContent)
						if len(matches) > 1 {
							fullThinking := strings.TrimSpace(matches[1])
							remaining := fullThinking[len(thinkingBuffer):]
							if remaining != "" {
								eventChan <- StreamEvent{
									Type: "thinking_chunk",
									Data: map[string]interface{}{
										"chunk":     remaining,
										"is_final":  true,
										"timestamp": time.Now().Format(time.RFC3339),
									},
								}
							}

							eventChan <- StreamEvent{
								Type: "thinking_end",
								Data: map[string]interface{}{
									"content":   fullThinking,
									"timestamp": time.Now().Format(time.RFC3339),
								},
							}
						}
						isInThinkingMode = false
					}

					if strings.Contains(accumulatedContent, "<message>") && !isInMessageMode {
						isInMessageMode = true
						hasStartedStreaming = true
						eventChan <- StreamEvent{
							Type:      "stream_start",
							Timestamp: time.Now().Format(time.RFC3339),
						}
					}

					if isInMessageMode && !strings.Contains(accumulatedContent, "</message>") {
						messageChunk := chunk
						messageChunk = regexp.MustCompile(`<thinking>[\s\S]*?</thinking>`).ReplaceAllString(messageChunk, "")
						messageChunk = strings.ReplaceAll(messageChunk, "</thinking>", "")
						messageChunk = strings.ReplaceAll(messageChunk, "<message>", "")
						messageChunk = strings.TrimSpace(messageChunk)

						if messageChunk != "" {
							messageBuffer += messageChunk
							eventChan <- StreamEvent{
								Type: "message_chunk",
								Data: map[string]interface{}{
									"chunk":     messageChunk,
									"is_final":  false,
									"timestamp": time.Now().Format(time.RFC3339),
								},
							}
						}
						continue
					}

					if strings.Contains(accumulatedContent, "</message>") && isInMessageMode {
						re := regexp.MustCompile(`<message>([\s\S]*?)</message>`)
						matches := re.FindStringSubmatch(accumulatedContent)
						if len(matches) > 1 {
							fullMessage := strings.TrimSpace(matches[1])
							remaining := fullMessage[len(messageBuffer):]
							if remaining != "" {
								eventChan <- StreamEvent{
									Type: "message_chunk",
									Data: map[string]interface{}{
										"chunk":     remaining,
										"is_final":  true,
										"timestamp": time.Now().Format(time.RFC3339),
									},
								}
							}
						}
						isInMessageMode = false
					}
				} else {
					if !hasStartedStreaming {
						hasStartedStreaming = true
						eventChan <- StreamEvent{
							Type:      "stream_start",
							Timestamp: time.Now().Format(time.RFC3339),
						}
					}

					eventChan <- StreamEvent{
						Type: "message_chunk",
						Data: map[string]interface{}{
							"chunk":     chunk,
							"is_final":  false,
							"timestamp": time.Now().Format(time.RFC3339),
						},
					}
				}

			case err := <-errChan:
				if err != nil {
					return err
				}
			}
		}

	StreamDone:
		response := &types.Message{
			Role:    "assistant",
			Content: accumulatedContent,
		}

		response = a.parseManualToolCalls(response)
		state.Messages = append(state.Messages, *response)

		if len(response.ToolCalls) == 0 {
			if hasStartedStreaming {
				eventChan <- StreamEvent{
					Type: "message_chunk",
					Data: map[string]interface{}{
						"chunk":     "",
						"is_final":  true,
						"timestamp": time.Now().Format(time.RFC3339),
					},
				}
			}

			eventChan <- StreamEvent{
				Type:      "stream_end",
				Timestamp: time.Now().Format(time.RFC3339),
			}

			lastMessage := response.Content
			state.Message = &lastMessage
			break
		}

		eventChan <- StreamEvent{
			Type: "node_execution",
			Data: map[string]interface{}{
				"node":       "agent_planning",
				"step":       stepCount,
				"tool_calls": response.ToolCalls,
				"timestamp":  time.Now().Format(time.RFC3339),
			},
		}
		stepCount++

		for _, tc := range response.ToolCalls {
			eventChan <- StreamEvent{
				Type: "node_execution",
				Data: map[string]interface{}{
					"node":      "tool_execution_start",
					"step":      stepCount,
					"tool_name": tc.Name,
					"tool_args": tc.Args,
					"timestamp": time.Now().Format(time.RFC3339),
				},
			}
			stepCount++
		}

		toolMessages, err := a.executeTools(requestID, ctx, response.ToolCalls)
		if err != nil {
			return err
		}

		for i, toolMsg := range toolMessages {
			eventChan <- StreamEvent{
				Type: "node_execution",
				Data: map[string]interface{}{
					"node":        "tool_execution_end",
					"step":        stepCount,
					"tool_name":   response.ToolCalls[i].Name,
					"tool_result": toolMsg.Content,
					"timestamp":   time.Now().Format(time.RFC3339),
				},
			}
			stepCount++
		}

		state.Messages = append(state.Messages, toolMessages...)
	}

	eventChan <- StreamEvent{
		Type: "done",
		Data: map[string]interface{}{
			"done":        true,
			"total_steps": stepCount,
			"timestamp":   time.Now().Format(time.RFC3339),
		},
	}

	return nil
}

func (a *LangChainAgent) buildMessages(requestID string, state *types.AgentState) []llms.MessageContent {
	var messages []llms.MessageContent

	// Add system prompt only once at the beginning
	if a.systemPrompt != nil && len(state.Messages) == 0 {
		messages = append(messages, llms.TextParts(llms.ChatMessageTypeSystem, *a.systemPrompt))
	}

	if !a.supportsTools && len(state.Messages) == 0 {
		reactPrompt := a.buildReactPrompt()
		messages = append(messages, llms.TextParts(llms.ChatMessageTypeSystem, reactPrompt))
	}

	// Implement sliding window to prevent context overflow
	// Keep only the last N messages to stay within context limits
	maxHistoryMessages := 4 // Default: keep last 4 messages (very conservative for local models)
	if a.llmClient.Config != nil && a.llmClient.Config.MaxContextMessages != nil {
		maxHistoryMessages = *a.llmClient.Config.MaxContextMessages
	}

	startIdx := 0
	if len(state.Messages) > maxHistoryMessages {
		startIdx = len(state.Messages) - maxHistoryMessages
		utils.VerbosePrintf("[%s]      ⚠️  Trimming message history: keeping last %d of %d messages\n", requestID, maxHistoryMessages, len(state.Messages))
	}

	for _, msg := range state.Messages[startIdx:] {
		var msgType llms.ChatMessageType
		switch msg.Role {
		case "user":
			msgType = llms.ChatMessageTypeHuman
		case "assistant":
			msgType = llms.ChatMessageTypeAI
		case "system":
			msgType = llms.ChatMessageTypeSystem
		default:
			msgType = llms.ChatMessageTypeGeneric
		}
		messages = append(messages, llms.TextParts(msgType, msg.Content))
	}

	messages = append(messages, llms.TextParts(llms.ChatMessageTypeHuman, state.Input))

	return messages
}

func (a *LangChainAgent) buildReactPrompt() string {
	var toolDescriptions []string
	for _, tool := range a.toolDefs {
		// Build parameter schema with required fields
		paramsJSON, _ := json.Marshal(tool.Parameters)

		// Build human-readable params list
		var params []string
		for key, prop := range tool.Parameters.Properties {
			isRequired := false
			for _, req := range tool.Parameters.Required {
				if req == key {
					isRequired = true
					break
				}
			}
			reqMark := ""
			if isRequired {
				reqMark = " [REQUIRED]"
			}
			params = append(params, fmt.Sprintf("%s: %s%s", key, prop.Type, reqMark))
		}

		desc := fmt.Sprintf("- %s(%s): %s\n  Schema: %s",
			tool.Name,
			strings.Join(params, ", "),
			tool.Description,
			string(paramsJSON))
		toolDescriptions = append(toolDescriptions, desc)
	}

	if a.provider == "llama_cpp" {
		return fmt.Sprintf(`You are a helpful AI assistant with access to these tools:
%s

RULES - READ CAREFULLY:
1. BE CONCISE. Keep <thinking> under 100 words.
2. Use tools when needed. Format: <thinking>brief reason</thinking><message>{"tool_name":"name","tool_args":{...}}</message>
3. For final answers: <thinking>brief summary</thinking><message>your answer</message>
4. NO extra text outside tags. NO repetition.
5. IMPORTANT: When using tools, carefully check ALL required parameters. Provide EXACTLY the parameters needed - no missing, no extra. Match parameter names and types precisely.
6. CRITICAL: After tool execution, CHECK THE RESULT. If result is SUCCESS (no "error" field), USE IT IMMEDIATELY in your final answer. DO NOT retry successful tools. Only retry if result contains "error" field (max 3 retries). If you get valid data, STOP calling tools and answer the user.

EXAMPLE (tool call):
<thinking>Need weather data for coordinates.</thinking>
<message>{"tool_name":"getWeather","tool_args":{"latitude":-7.7,"longitude":109.0}}</message>

EXAMPLE (final answer):
<thinking>Tool returned weather data.</thinking>
<message>The weather is sunny, 28°C.</message>

BE BRIEF. NO REPETITION.`, strings.Join(toolDescriptions, "\n"))
	}

	return fmt.Sprintf(`You have access to the following tools:

%s

To use a tool, you MUST respond with ONLY a JSON object in this EXACT format:
{"tool_name": "name_of_tool", "tool_args": {"param1": "value1", "param2": value2}}

Do NOT add any explanation before or after the JSON. Just output the JSON.

If you don't need a tool, respond normally to the user's question.`, strings.Join(toolDescriptions, "\n"))
}

func (a *LangChainAgent) parseManualToolCalls(response *types.Message) *types.Message {
	content := strings.TrimSpace(response.Content)

	re1 := regexp.MustCompile(`\{\s*"tool_name"\s*:\s*"([^"]+)"\s*,\s*"tool_args"\s*:\s*(\{[^}]*\})\s*\}`)
	matches := re1.FindStringSubmatch(content)

	if matches == nil {
		re2 := regexp.MustCompile(`to=(?:tool\.)?(\w+)\s+json\s*\n?\s*(\{[^}]+\})`)
		matches = re2.FindStringSubmatch(content)
	}

	// Pattern 3: to=tool.toolname code<|message|>{...}
	if matches == nil {
		re3 := regexp.MustCompile(`to=tool\.(\w+)\s+code<\|message\|>(\{[^}]+\})`)
		matches = re3.FindStringSubmatch(content)
	}

	// Pattern 4: <message>{...} after tool mention
	if matches == nil {
		re4 := regexp.MustCompile(`<message>(\{[^}]+\})`)
		jsonMatch := re4.FindStringSubmatch(content)
		if jsonMatch != nil {
			// Try to find tool name from context
			toolNameRe := regexp.MustCompile(`(?:to=tool\.|dns_lookup|tool_name["\s:]+)(\w+)`)
			toolMatch := toolNameRe.FindStringSubmatch(content)
			if toolMatch != nil {
				matches = []string{"", toolMatch[1], jsonMatch[1]}
			}
		}
	}

	if matches != nil && len(matches) >= 3 {
		toolName := matches[1]
		argsStr := matches[2]

		var args map[string]interface{}
		if err := json.Unmarshal([]byte(argsStr), &args); err == nil {
			for key, val := range args {
				if strVal, ok := val.(string); ok {
					if num, err := strconv.ParseFloat(strVal, 64); err == nil {
						args[key] = num
					}
				}
			}

			response.ToolCalls = []types.ToolCall{
				{
					ID:   fmt.Sprintf("manual_%d", time.Now().UnixNano()),
					Name: toolName,
					Args: args,
					Type: "tool_call",
				},
			}
		}
	}

	return response
}

func (a *LangChainAgent) executeTools(requestID string, ctx context.Context, toolCalls []types.ToolCall) ([]types.Message, error) {
	var toolMessages []types.Message

	for idx, call := range toolCalls {
		utils.VerbosePrintf("[%s]         [%d/%d] Executing: %s\n", requestID, idx+1, len(toolCalls), call.Name)
		var result interface{}
		var err error

		for _, serverURL := range a.mcpServers {
			result, err = mcp.InvokeTool(serverURL, call.Name, call.Args)
			if err == nil {
				utils.VerbosePrintf("[%s]            ✅ Success from %s\n", requestID, serverURL)
				break
			} else {
				utils.VerbosePrintf("[%s]            ⚠️  Failed from %s: %v\n", requestID, serverURL, err)
			}
		}

		if err != nil {
			utils.VerbosePrintf("[%s]            ❌ All servers failed for %s\n", requestID, call.Name)
			return nil, err
		}

		resultJSON, _ := json.Marshal(result)

		toolMessages = append(toolMessages, types.Message{
			Role:       "tool",
			ToolCallID: call.ID,
			Name:       call.Name,
			Content:    string(resultJSON),
		})
	}

	return toolMessages, nil
}

type StreamEvent struct {
	Type      string                 `json:"type"`
	Timestamp string                 `json:"timestamp,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}
