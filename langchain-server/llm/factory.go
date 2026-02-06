package llm

import (
	"context"
	"fmt"

	"langchain-mcp-api/types"

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

func CreateLangChainLLM(credential types.RequestChatCredential) (*LangChainClient, error) {
	fmt.Println("\n🔧 [LLM] Creating LangChain LLM client...")
	provider := credential.Provider

	model := DefaultModelsLangChain[provider]
	if credential.Model != nil {
		model = *credential.Model
	}
	fmt.Printf("   Provider: %s\n", provider)
	fmt.Printf("   Model: %s\n", model)

	client := &LangChainClient{
		Provider: provider,
		Model:    model,
		URL:      credential.URL,
		Config:   credential.Set,
	}

	if credential.Set != nil {
		fmt.Println("   Configuration:")
		if credential.Set.Temperature != nil {
			fmt.Printf("      Temperature: %.2f\n", *credential.Set.Temperature)
		}
		if credential.Set.MaxTokens != nil {
			fmt.Printf("      MaxTokens: %d\n", *credential.Set.MaxTokens)
		}
		if credential.Set.TopP != nil {
			fmt.Printf("      TopP: %.2f\n", *credential.Set.TopP)
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

		fmt.Printf("   Llama.cpp BaseURL: %s\n", *credential.URL)
		llmInstance, err = openai.New(
			openai.WithToken("llama_cpp"),
			openai.WithModel(model),
			openai.WithBaseURL(*credential.URL),
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
		fmt.Printf("   ❌ Failed to create LLM: %v\n", err)
		return nil, err
	}

	fmt.Println("   LLM client created successfully")
	client.LLM = llmInstance
	fmt.Printf("\n [LLM] LangChain LLM client created (provider: %s, model: %s)\n", client.Provider, client.Model)
	return client, nil
}

func (c *LangChainClient) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (string, error) {
	fmt.Printf("\n [LLM] GenerateContent called (provider: %s, model: %s)\n", c.Provider, c.Model)
	fmt.Printf("   Messages: %d\n", len(messages))

	// Build call options from config
	callOpts := c.buildCallOptions()
	callOpts = append(callOpts, options...)
	fmt.Printf("   CallOptions: %d\n", len(callOpts))

	result, err := c.LLM.GenerateContent(ctx, messages, callOpts...)
	if err != nil {
		fmt.Printf("   ❌ GenerateContent error: %v\n", err)
		fmt.Printf("   Provider: %s, Model: %s\n", c.Provider, c.Model)
		if c.URL != nil {
			fmt.Printf("   Configured URL: %s\n", *c.URL)
		}
		return "", err
	}

	if len(result.Choices) == 0 {
		fmt.Println("   No response choices from LLM")
		return "", fmt.Errorf("no response from LLM")
	}

	content := result.Choices[0].Content
	fmt.Printf("   Response received (%d chars)\n", len(content))
	return content, nil
}

func (c *LangChainClient) StreamGenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (<-chan string, <-chan error) {
	fmt.Printf("\n [LLM] StreamGenerateContent called (provider: %s, model: %s)\n", c.Provider, c.Model)
	fmt.Printf("   Messages: %d\n", len(messages))

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
				fmt.Printf("   📦 Received %d chunks...\n", chunkCount)
			}
			contentChan <- string(chunk)
			return nil
		}))

		_, err := c.LLM.GenerateContent(ctx, messages, callOpts...)
		if err != nil {
			fmt.Printf("   ❌ Streaming error: %v\n", err)
			errChan <- err
		} else {
			fmt.Printf("   ✅ Streaming completed (%d total chunks)\n", chunkCount)
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
