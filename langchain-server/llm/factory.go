package llm

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"langchain-mcp-api/types"
	"langchain-mcp-api/utils"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

var DefaultModelsLangChain = map[string]string{
	"openai":     "gpt-4o-mini",
	"claude":     "claude-3-5-sonnet-20241022",
	"openrouter": "anthropic/claude-3.5-sonnet",
	"ollama":     "llama3.2",
	"llama_cpp":  "gpt-oss-20b.gguf",
	"vllm":       "meta-llama/Llama-3.2-3B-Instruct",
}

type LangChainClient struct {
	LLM           llms.Model
	Provider      string
	Model         string
	URL           *string
	Config        *types.SetLLM
	SupportsTools bool
}

func CreateLangChainLLM(requestID string, credential types.RequestChatCredential) (*LangChainClient, error) {
	utils.VerbosePrintf("\n[%s]🔧 [LLM] Creating LangChain LLM client...\n", requestID)
	provider := credential.Provider

	model := DefaultModelsLangChain[provider]
	if credential.Model != nil {
		model = *credential.Model
	}
	utils.VerbosePrintf("[%s]   Provider: %s\n", requestID, provider)
	utils.VerbosePrintf("[%s]   Model: %s\n", requestID, model)

	client := &LangChainClient{
		Provider: provider,
		Model:    model,
		URL:      credential.URL,
		Config:   credential.Set,
	}

	if credential.Set != nil {
		utils.VerbosePrintf("[%s]   Configuration:\n", requestID)
		if credential.Set.Temperature != nil {
			utils.VerbosePrintf("[%s]      Temperature: %.2f\n", requestID, *credential.Set.Temperature)
		}
		if credential.Set.MaxTokens != nil {
			utils.VerbosePrintf("[%s]      MaxTokens: %d\n", requestID, *credential.Set.MaxTokens)
		}
		if credential.Set.TopP != nil {
			utils.VerbosePrintf("[%s]      TopP: %.2f\n", requestID, *credential.Set.TopP)
		}
	}

	toolCallingProviders := []string{"openai", "claude", "openrouter"}
	client.SupportsTools = containsString(toolCallingProviders, provider)

	var llmInstance llms.Model
	var err error

	switch provider {
	case "openai":
		if credential.APIKey == nil {
			return nil, types.NewErrorRequest("OpenAI API key is required", 401)
		}

		llmInstance, err = openai.New(
			openai.WithToken(*credential.APIKey),
			openai.WithModel(model),
		)

	case "claude":
		if credential.APIKey == nil {
			return nil, types.NewErrorRequest("Claude API key is required", 401)
		}

		llmInstance, err = anthropic.New(
			anthropic.WithToken(*credential.APIKey),
			anthropic.WithModel(model),
		)

	case "openrouter":
		if credential.APIKey == nil {
			return nil, types.NewErrorRequest("OpenRouter API key is required", 401)
		}

		llmInstance, err = openai.New(
			openai.WithToken(*credential.APIKey),
			openai.WithModel(model),
			openai.WithBaseURL("https://openrouter.ai/api/v1"),
		)

	case "ollama":
		if credential.URL == nil {
			return nil, types.NewErrorRequest("Ollama URL is required", 400)
		}

		llmInstance, err = ollama.New(
			ollama.WithModel(model),
			ollama.WithServerURL(*credential.URL),
		)

	case "llama_cpp":
		if credential.URL == nil {
			return nil, types.NewErrorRequest("Llama.cpp URL is required", 400)
		}

		utils.VerbosePrintf("[%s]   Llama.cpp BaseURL: %s\n", requestID, *credential.URL)

		// Create custom HTTP client with longer timeout and TLS skip verify
		httpClient := &http.Client{
			Timeout: 300 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		}
		utils.VerbosePrintf("[%s]   Using custom HTTP client (timeout: 300s, TLS skip verify: true)\n", requestID)

		llmInstance, err = openai.New(
			openai.WithToken("llama_cpp"),
			openai.WithModel(model),
			openai.WithBaseURL(*credential.URL),
			openai.WithHTTPClient(httpClient),
		)

	case "vllm":
		if credential.URL == nil {
			return nil, types.NewErrorRequest("vLLM URL is required", 400)
		}

		llmInstance, err = openai.New(
			openai.WithToken("vllm"),
			openai.WithModel(model),
			openai.WithBaseURL(*credential.URL),
		)

	default:
		return nil, types.NewErrorRequest(fmt.Sprintf("Unsupported provider: %s", provider), 404)
	}

	if err != nil {
		utils.VerbosePrintf("[%s]   ❌ Failed to create LLM: %v\n", requestID, err)
		return nil, err
	}

	utils.VerbosePrintf("[%s]   LLM client created successfully\n", requestID)
	client.LLM = llmInstance
	utils.VerbosePrintf("\n[%s] [LLM] LangChain LLM client created (provider: %s, model: %s)\n", requestID, client.Provider, client.Model)
	return client, nil
}

func (c *LangChainClient) GenerateContent(requestID string, ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (string, error) {
	utils.VerbosePrintf("\n[%s] [LLM] GenerateContent called (provider: %s, model: %s)\n", requestID, c.Provider, c.Model)
	utils.VerbosePrintf("[%s]   Messages: %d\n", requestID, len(messages))

	// Build call options from config
	callOpts := c.buildCallOptions()
	callOpts = append(callOpts, options...)
	utils.VerbosePrintf("[%s]   CallOptions: %d\n", requestID, len(callOpts))

	result, err := c.LLM.GenerateContent(ctx, messages, callOpts...)
	if err != nil {
		utils.VerbosePrintf("[%s]   ❌ GenerateContent error: %v\n", requestID, err)
		utils.VerbosePrintf("[%s]   Provider: %s, Model: %s\n", requestID, c.Provider, c.Model)
		if c.URL != nil {
			utils.VerbosePrintf("[%s]   Configured URL: %s\n", requestID, *c.URL)
		}
		return "", err
	}

	if len(result.Choices) == 0 {
		utils.VerbosePrintf("[%s]   No response choices from LLM\n", requestID)
		return "", fmt.Errorf("no response from LLM")
	}

	content := result.Choices[0].Content
	utils.VerbosePrintf("[%s]   Response received (%d chars)\n", requestID, len(content))
	return content, nil
}

func (c *LangChainClient) StreamGenerateContent(requestID string, ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (<-chan string, <-chan error) {
	utils.VerbosePrintf("\n[%s] [LLM] StreamGenerateContent called (provider: %s, model: %s)\n", requestID, c.Provider, c.Model)
	utils.VerbosePrintf("[%s]   Messages: %d\n", requestID, len(messages))

	contentChan := make(chan string, 100)
	errChan := make(chan error, 1)

	go func() {
		defer close(contentChan)
		defer close(errChan)

		// Build call options from config
		callOpts := c.buildCallOptions()
		callOpts = append(callOpts, options...)

		chunkCount := 0
		callOpts = append(callOpts, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			chunkCount++
			if chunkCount%10 == 0 {
				utils.VerbosePrintf("[%s]   📦 Received %d chunks...\n", requestID, chunkCount)
			}
			contentChan <- string(chunk)
			return nil
		}))

		_, err := c.LLM.GenerateContent(ctx, messages, callOpts...)
		if err != nil {
			utils.VerbosePrintf("[%s]   ❌ Streaming error: %v\n", requestID, err)
			errChan <- err
		} else {
			utils.VerbosePrintf("[%s]   ✅ Streaming completed (%d total chunks)\n", requestID, chunkCount)
		}
	}()

	return contentChan, errChan
}

func (c *LangChainClient) buildCallOptions() []llms.CallOption {
	var opts []llms.CallOption

	if c.Config == nil {
		return opts
	}

	if c.Config.Temperature != nil {
		opts = append(opts, llms.WithTemperature(*c.Config.Temperature))
	}
	if c.Config.MaxTokens != nil {
		opts = append(opts, llms.WithMaxTokens(*c.Config.MaxTokens))
	}
	if c.Config.TopP != nil {
		opts = append(opts, llms.WithTopP(*c.Config.TopP))
	}
	if c.Config.FrequencyPenalty != nil {
		opts = append(opts, llms.WithFrequencyPenalty(*c.Config.FrequencyPenalty))
	}
	if c.Config.PresencePenalty != nil {
		opts = append(opts, llms.WithPresencePenalty(*c.Config.PresencePenalty))
	}
	if len(c.Config.Stop) > 0 {
		opts = append(opts, llms.WithStopWords(c.Config.Stop))
	}

	return opts
}

func containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
