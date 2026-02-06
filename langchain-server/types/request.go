package types

type RequestChatBody struct {
	Credential   RequestChatCredential `json:"credential"`
	SystemPrompt *string               `json:"system_prompt,omitempty"`
	Input        string                `json:"input"`
	Servers      []string              `json:"servers"`
}

type RequestChatCredential struct {
	Provider string  `json:"provider"`
	Model    *string `json:"model,omitempty"`
	URL      *string `json:"url,omitempty"`
	APIKey   *string `json:"api_key,omitempty"`
	Set      *SetLLM `json:"set,omitempty"`
}

type SetLLM struct {
	Temperature        *float64 `json:"temperature,omitempty"`
	MaxTokens          *int     `json:"max_tokens,omitempty"`
	TopP               *float64 `json:"top_p,omitempty"`
	FrequencyPenalty   *float64 `json:"frequency_penalty,omitempty"`
	PresencePenalty    *float64 `json:"presence_penalty,omitempty"`
	Stop               []string `json:"stop,omitempty"`
	Timeout            *int     `json:"timeout,omitempty"`
	MaxRetries         *int     `json:"max_retries,omitempty"`
	MaxContextMessages *int     `json:"max_context_messages,omitempty"` // Limit message history to prevent context overflow
}

type LLMPublicProvider string

const (
	ProviderOpenAI     LLMPublicProvider = "openai"
	ProviderClaude     LLMPublicProvider = "claude"
	ProviderOpenRouter LLMPublicProvider = "openrouter"
)

type LLMLocalProvider string

const (
	ProviderOllama   LLMLocalProvider = "ollama"
	ProviderLlamaCPP LLMLocalProvider = "llama_cpp"
	ProviderVLLM     LLMLocalProvider = "vllm"
)
