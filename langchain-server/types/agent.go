package types

type ToolCall struct {
	ID   string                 `json:"id"`
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args"`
	Type string                 `json:"type"`
}

type TokenUsage struct {
	PromptTokens     int `json:"promptTokens"`
	CompletionTokens int `json:"completionTokens"`
	TotalTokens      int `json:"totalTokens"`
}

type UsageMetadata struct {
	OutputTokens       int                 `json:"output_tokens"`
	InputTokens        int                 `json:"input_tokens"`
	TotalTokens        int                 `json:"total_tokens"`
	InputTokenDetails  *InputTokenDetails  `json:"input_token_details,omitempty"`
	OutputTokenDetails *OutputTokenDetails `json:"output_token_details,omitempty"`
}

type InputTokenDetails struct {
	Audio     int `json:"audio"`
	CacheRead int `json:"cache_read"`
}

type OutputTokenDetails struct {
	Audio     int `json:"audio"`
	Reasoning int `json:"reasoning"`
}

type ResponseMetadata struct {
	TokenUsage        TokenUsage  `json:"tokenUsage"`
	FinishReason      string      `json:"finish_reason"`
	ModelProvider     string      `json:"model_provider"`
	ModelName         string      `json:"model_name"`
	Usage             interface{} `json:"usage"`
	SystemFingerprint string      `json:"system_fingerprint"`
}

type Message struct {
	Role       string            `json:"role"`
	Content    string            `json:"content"`
	ToolCalls  []ToolCall        `json:"tool_calls,omitempty"`
	ToolCallID string            `json:"tool_call_id,omitempty"`
	Name       string            `json:"name,omitempty"`
	ID         string            `json:"id,omitempty"`
	Metadata   *ResponseMetadata `json:"response_metadata,omitempty"`
	UsageData  *UsageMetadata    `json:"usage_metadata,omitempty"`
}

type AgentState struct {
	Input    string    `json:"input"`
	Messages []Message `json:"messages"`
	Message  *string   `json:"message,omitempty"`
}

type ChatResponse struct {
	Messages []Message `json:"messages"`
	Message  string    `json:"message"`
}
