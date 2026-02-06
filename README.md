<div align="center">

# 🤖 LangChain MCP API

**High-Performance Go Implementation with Multi-Provider LLM Support**

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Fiber](https://img.shields.io/badge/Fiber-v3-00ACD7?style=flat&logo=go)](https://gofiber.io)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://hub.docker.com)

*Universal LangChain server supporting OpenAI, Claude, Ollama, Llama.cpp, and more with MCP tools integration*

[Features](#-features) • [Quick Start](#-quick-start) • [API Documentation](#-api-documentation) • [Examples](#-examples) • [Docker](#-docker)

</div>

---

## ✨ Features

<table>
<tr>
<td width="50%">

### 🚀 **Multiple LLM Providers**
- **OpenAI** (GPT-4, GPT-3.5, GPT-4o)
- **Claude** (Anthropic)
- **OpenRouter** (100+ models)
- **Ollama** (Local models)
- **Llama.cpp** (GGUF models)
- **vLLM** (High-performance inference)

</td>
<td width="50%">

### 🔧 **Advanced Features**
- **MCP Tools** - Dynamic tool loading
- **Streaming** - Real-time SSE responses
- **Agent System** - Autonomous task execution
- **Context Management** - Smart history trimming
- **Verbose Logging** - Detailed execution traces

</td>
</tr>
</table>

---

## 🚀 Quick Start

### Prerequisites
- Go 1.25 or higher
- MCP server (optional, for tools)

### Installation

```bash
# Clone the repository
git clone https://github.com/jefripunza/langchain-server.git
cd langchain-server/langchain-server

# Install dependencies
go mod tidy

# Run the server
go run main.go
```

Server will start at `http://localhost:6000` 🎉

---

## 📡 API Documentation

### Base URL
```
http://localhost:6000
```

### Endpoints

#### 1️⃣ **Health Check**

```http
GET /
```

**Response:**
```json
{
  "message": "🤖 LangChain MCP API is running",
  "version": "1.0.0"
}
```

---

#### 2️⃣ **Chat (Non-Streaming)**

```http
POST /chat
```

**Request Body:**
```json
{
  "credential": {
    "provider": "openai",
    "api_key": "sk-...",
    "model": "gpt-4o-mini",
    "set": {
      "temperature": 0.7,
      "max_tokens": 1000,
      "max_context_messages": 4
    }
  },
  "system_prompt": "You are a helpful assistant",
  "input": "What is the weather in Jakarta?",
  "servers": ["http://localhost:3001"]
}
```

**Response:**
```json
{
  "input": "What is the weather in Jakarta?",
  "messages": [
    {
      "role": "assistant",
      "content": "I'll check the weather for you...",
      "tool_calls": [...]
    },
    {
      "role": "tool",
      "content": "{\"temperature\": 28, \"condition\": \"sunny\"}",
      "name": "getWeather"
    }
  ],
  "message": "The weather in Jakarta is sunny with a temperature of 28°C."
}
```

---

#### 3️⃣ **Chat Stream (SSE)**

```http
POST /chat/stream
```

**Request Body:** *(same as `/chat`)*

**Response:** Server-Sent Events stream

```
data: {"type":"start","timestamp":"2024-02-04T09:00:00Z","input":"What is the weather?"}

data: {"type":"servers_checked","available_servers":["http://localhost:3001"],"total_servers":1}

data: {"type":"thinking_start","timestamp":"2024-02-04T09:00:01Z"}

data: {"type":"thinking_chunk","chunk":"I need to check the weather...","is_final":false}

data: {"type":"message_start","timestamp":"2024-02-04T09:00:02Z"}

data: {"type":"message_chunk","chunk":"The weather is ","is_final":false}

data: {"type":"message_chunk","chunk":"sunny, 28°C","is_final":true}

data: {"type":"done","done":true,"total_steps":3,"timestamp":"2024-02-04T09:00:03Z"}
```

---

## 🎯 Examples

### Example 1: OpenAI with Tools

```bash
curl -X POST http://localhost:6000/chat \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "openai",
      "api_key": "sk-proj-...",
      "model": "gpt-4o-mini"
    },
    "input": "What is 25 + 37?",
    "servers": ["http://localhost:3001"]
  }'
```

### Example 2: Claude Streaming

```bash
curl -N -X POST http://localhost:6000/chat/stream \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "claude",
      "api_key": "sk-ant-...",
      "model": "claude-3-5-sonnet-20241022"
    },
    "input": "Explain quantum computing in simple terms",
    "servers": []
  }'
```

### Example 3: Local Llama.cpp

```bash
curl -X POST http://localhost:6000/chat \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "llama_cpp",
      "url": "http://localhost:8080",
      "model": "gpt-oss-20b.gguf",
      "set": {
        "temperature": 0.7,
        "max_context_messages": 4
      }
    },
    "system_prompt": "You are a helpful assistant. Be concise.",
    "input": "What is the capital of France?",
    "servers": ["http://localhost:3001"]
  }'
```

### Example 4: Ollama with Custom Settings

```bash
curl -X POST http://localhost:6000/chat \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "ollama",
      "url": "http://localhost:11434",
      "model": "llama3.2",
      "set": {
        "temperature": 0.8,
        "max_tokens": 500
      }
    },
    "input": "Write a haiku about programming",
    "servers": []
  }'
```

---

## 🐳 Docker

### Build and Run

```bash
# Build image
docker build -t langchain-server .

# Run container
docker run -d \
  --name langchain-server \
  -p 6000:6000 \
  langchain-server
```

### Using Docker Compose

```yaml
version: '3.8'
services:
  langchain-server:
    build: .
    ports:
      - "6000:6000"
    environment:
      - PORT=6000
    restart: unless-stopped
```

---

## ⚙️ Configuration

### Provider Settings

| Provider | Required Fields | Optional Fields |
|----------|----------------|-----------------|
| **OpenAI** | `api_key`, `model` | `temperature`, `max_tokens`, `top_p` |
| **Claude** | `api_key`, `model` | `temperature`, `max_tokens`, `top_p` |
| **OpenRouter** | `api_key`, `model` | `temperature`, `max_tokens` |
| **Ollama** | `url`, `model` | `temperature` |
| **Llama.cpp** | `url`, `model` | `temperature`, `max_tokens` |
| **vLLM** | `url`, `model` | `temperature`, `max_tokens` |

### Advanced Settings

```json
{
  "set": {
    "temperature": 0.7,           // Creativity (0.0 - 2.0)
    "max_tokens": 1000,           // Max response length
    "top_p": 0.9,                 // Nucleus sampling
    "frequency_penalty": 0.0,     // Repetition penalty
    "presence_penalty": 0.0,      // Topic diversity
    "max_context_messages": 4     // History window size
  }
}
```

---

## 🔍 Verbose Logging

The server provides detailed execution logs:

```
📦 [AGENT] Creating LangChain Agent...
   Provider: llama_cpp
   Model: gpt-oss-20b.gguf
   ✅ Loaded 22 tools from MCP servers

🚀 [INVOKE] Starting agent invocation...
   Input: What is the weather?

   🔁 [ITERATION 1/10]
      📝 Built 2 messages for LLM
      🤖 Calling LLM...
      ✅ LLM Response (245 chars)
      🔧 Detected 1 tool call(s)
         1. getWeather({"lat": -7.7, "lon": 109.0})
      ⚙️  Executing tools...
         [1/1] Executing: getWeather
            ✅ Success from http://localhost:3001
      ✅ Tools executed successfully (1 results)

✅ [INVOKE] Agent invocation completed
```

---

## 🛠️ Development

### Project Structure

```
langchain-server/
├── main.go              # Entry point
├── agent/               # Agent logic
│   └── agent.go
├── handlers/            # HTTP handlers
│   └── chat.go
├── llm/                 # LLM factory
│   └── factory.go
├── mcp/                 # MCP tools loader
│   └── loader.go
├── types/               # Type definitions
│   ├── request.go
│   ├── agent.go
│   ├── tool.go
│   └── error.go
└── utils/               # Utilities
    └── utils.go
```

### Build for Production

```bash
# Build binary
go build -o langchain-server

# Run binary
./langchain-server
```

### Run Tests

```bash
go test ./...
```

---

## 🐛 Troubleshooting

<details>
<summary><b>Error: "No MCP servers available"</b></summary>

- Ensure MCP server is running
- Check server URL is correct
- Verify health endpoint: `curl http://localhost:3001/health`
</details>

<details>
<summary><b>Error: "Missing api key"</b></summary>

- Verify API key is set in request
- Check API key format is correct
- Ensure provider name matches
</details>

<details>
<summary><b>Error: "Context size exceeded"</b></summary>

- Reduce `max_context_messages` (default: 4)
- Use shorter system prompts
- Enable response truncation
</details>

<details>
<summary><b>Streaming not working</b></summary>

- Ensure client supports Server-Sent Events
- Check network/proxy settings
- Use `-N` flag with curl for streaming
</details>

---

## 📊 Performance

| Metric | Value |
|--------|-------|
| **Startup Time** | < 1s |
| **Memory Usage** | ~50MB (idle) |
| **Concurrent Requests** | 1000+ |
| **Response Time** | < 100ms (without LLM) |

---

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details

---

## 👥 Contributors

<a href="https://github.com/jefripunza">
  <img src="https://github.com/jefripunza.png" width="50" height="50" alt="Jefri Herdi Triyanto" style="border-radius: 50%;">
</a>

**Jefri Herdi Triyanto** ([@jefripunza](https://github.com/jefripunza))

---

## 🔗 Related Projects

- [MCP Server](../mcp-server) - Model Context Protocol server implementation
- [LangChain Go](https://github.com/tmc/langchaingo) - Official LangChain Go library

---

<div align="center">

**⭐ Star this repo if you find it useful!**

Made with ❤️ using [Go](https://golang.org) and [Fiber](https://gofiber.io)

</div>
