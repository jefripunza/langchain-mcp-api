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

// GenerationInfo represents detailed token usage from LLM response
type GenerationInfo struct {
	CompletionTokens                   int    `json:"CompletionTokens"`
	PromptTokens                       int    `json:"PromptTokens"`
	TotalTokens                        int    `json:"TotalTokens"`
	CompletionAcceptedPredictionTokens int    `json:"CompletionAcceptedPredictionTokens"`
	CompletionAudioTokens              int    `json:"CompletionAudioTokens"`
	CompletionReasoningTokens          int    `json:"CompletionReasoningTokens"`
	CompletionRejectedPredictionTokens int    `json:"CompletionRejectedPredictionTokens"`
	PromptAudioTokens                  int    `json:"PromptAudioTokens"`
	PromptCachedTokens                 int    `json:"PromptCachedTokens"`
	ReasoningTokens                    int    `json:"ReasoningTokens"`
	ThinkingTokens                     int    `json:"ThinkingTokens"`
	ThinkingContent                    string `json:"ThinkingContent,omitempty"`
	ReasoningContent                   string `json:"ReasoningContent,omitempty"`
}

type UsageMetadata struct {
	OutputTokens                       int                 `json:"output_tokens"`
	InputTokens                        int                 `json:"input_tokens"`
	TotalTokens                        int                 `json:"total_tokens"`
	InputTokenDetails                  *InputTokenDetails  `json:"input_token_details,omitempty"`
	OutputTokenDetails                 *OutputTokenDetails `json:"output_token_details,omitempty"`
	CompletionAcceptedPredictionTokens int                 `json:"completion_accepted_prediction_tokens,omitempty"`
	CompletionAudioTokens              int                 `json:"completion_audio_tokens,omitempty"`
	CompletionReasoningTokens          int                 `json:"completion_reasoning_tokens,omitempty"`
	CompletionRejectedPredictionTokens int                 `json:"completion_rejected_prediction_tokens,omitempty"`
	PromptAudioTokens                  int                 `json:"prompt_audio_tokens,omitempty"`
	PromptCachedTokens                 int                 `json:"prompt_cached_tokens,omitempty"`
	ReasoningTokens                    int                 `json:"reasoning_tokens,omitempty"`
	ThinkingTokens                     int                 `json:"thinking_tokens,omitempty"`
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
	Messages         []Message         `json:"messages"` // ada
	Message          string            `json:"message"`  // ada
	Metadata         *ResponseMetadata `json:"metadata,omitempty"`
	UsageMetadata    *UsageMetadata    `json:"usage_metadata,omitempty"`
	TotalTokens      int               `json:"total_tokens,omitempty"`
	InputTokens      int               `json:"input_tokens,omitempty"`
	OutputTokens     int               `json:"output_tokens,omitempty"`
	ModelProvider    string            `json:"model_provider,omitempty"` // ada
	ModelName        string            `json:"model_name,omitempty"`     // ada
	FinishReason     string            `json:"finish_reason,omitempty"`
	TotalIterations  int               `json:"total_iterations,omitempty"`   // ada
	ToolCallsCount   int               `json:"tool_calls_count,omitempty"`   // ada
	ExecutionTimeMs  int64             `json:"execution_time_ms,omitempty"`  // ada
	ExecutionTimeSec float64           `json:"execution_time_sec,omitempty"` // ada
	TokensPerSecond  float64           `json:"tokens_per_second,omitempty"`
}
