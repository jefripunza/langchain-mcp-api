# LangChain Server - Go Implementation

Implementasi Go dari LangChain Server yang mendukung multiple LLM providers dan MCP (Model Context Protocol) tools.

## 🚀 Fitur

- **Multiple LLM Providers:**
  - OpenAI (GPT-4, GPT-3.5, dll)
  - Claude (Anthropic)
  - OpenRouter
  - Ollama (local)
  - Llama.cpp (local)
  - vLLM (local)

- **MCP Tools Integration:**
  - Dynamic tool loading dari MCP servers
  - Tool execution dengan error handling
  - Health check untuk MCP servers

- **Streaming Support:**
  - Server-Sent Events (SSE) untuk real-time streaming
  - Support untuk thinking mode (llama.cpp)
  - Progress tracking untuk tool execution

- **Agent System:**
  - State graph untuk agent workflow
  - Native tool calling untuk OpenAI/Claude/OpenRouter
  - Manual tool calling untuk Ollama/Llama.cpp/vLLM
  - Automatic tool parsing dan execution

## 📁 Struktur Project

```
langchain-server/
├── main.go                 # Entry point aplikasi
├── types/                  # Type definitions
│   ├── request.go         # Request types
│   ├── agent.go           # Agent types
│   ├── tool.go            # Tool types
│   └── error.go           # Error handling
├── utils/                  # Utilities
│   └── utils.go           # Helper functions
├── llm/                    # LLM factory
│   └── factory.go         # LLM client implementation
├── mcp/                    # MCP tools
│   └── loader.go          # MCP tools loader
├── agent/                  # Agent logic
│   └── agent.go           # Agent implementation
└── handlers/               # HTTP handlers
    └── chat.go            # Chat endpoints
```

## 🛠️ Installation

1. **Clone repository:**
```bash
cd /Users/jefripunza/Documents/Projects/RnD/bun-mcp-api/langchain-server
```

2. **Install dependencies:**
```bash
go mod tidy
```

3. **Build aplikasi:**
```bash
go build -o langchain-server
```

## 🏃 Running

**Development mode:**
```bash
go run main.go
```

**Production mode:**
```bash
./langchain-server
```

Server akan berjalan di `http://localhost:6000`

## 📡 API Endpoints

### 1. Health Check
```bash
GET /
```

Response:
```json
{
  "message": "🤖 LangChain Server is running",
  "version": "1.0.0"
}
```

### 2. Chat (Non-Streaming)
```bash
POST /chat
```

Request Body:
```json
{
  "credential": {
    "provider": "openai",
    "api_key": "sk-...",
    "model": "gpt-4o-mini",
    "set": {
      "temperature": 0.7,
      "max_tokens": 1000
    }
  },
  "system_prompt": "You are a helpful assistant",
  "input": "What is the weather in Jakarta?",
  "servers": ["http://localhost:3001"]
}
```

Response:
```json
{
  "messages": [...],
  "message": "The weather in Jakarta is..."
}
```

### 3. Chat Stream (SSE)
```bash
POST /chat/stream
```

Request Body: (sama dengan `/chat`)

Response: Server-Sent Events
```
data: {"type":"start","timestamp":"2024-02-04T07:27:00Z","input":"..."}
data: {"type":"servers_checked","available_servers":[...],"total_servers":1}
data: {"type":"thinking_start","timestamp":"..."}
data: {"type":"thinking_chunk","chunk":"...","is_final":false}
data: {"type":"message_chunk","chunk":"...","is_final":false}
data: {"type":"done","done":true,"total_steps":5}
```

## 🔧 Configuration

### Environment Variables (Optional)

```bash
# Server port
PORT=6000

# Default timeouts
HTTP_TIMEOUT=30000
```

### Provider Configuration

**OpenAI:**
```json
{
  "provider": "openai",
  "api_key": "sk-...",
  "model": "gpt-4o-mini"
}
```

**Claude:**
```json
{
  "provider": "claude",
  "api_key": "sk-ant-...",
  "model": "claude-3-5-sonnet-20241022"
}
```

**Ollama (Local):**
```json
{
  "provider": "ollama",
  "url": "http://localhost:11434",
  "model": "llama3.2"
}
```

**Llama.cpp (Local):**
```json
{
  "provider": "llama_cpp",
  "url": "http://localhost:8080",
  "model": "gpt-oss-20b.gguf"
}
```

## 📝 Example Usage

### Curl Example

```bash
curl -X POST http://localhost:6000/chat \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "openai",
      "api_key": "sk-...",
      "model": "gpt-4o-mini"
    },
    "input": "Hello, how are you?",
    "servers": ["http://localhost:3001"]
  }'
```

### Streaming Example

```bash
curl -X POST http://localhost:6000/chat/stream \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "llama_cpp",
      "url": "http://localhost:8080"
    },
    "input": "Explain quantum computing",
    "servers": ["http://localhost:3001"]
  }'
```

## 🧪 Testing

```bash
go test ./...
```

## 📦 Dependencies

- `github.com/gofiber/fiber/v3` - Web framework
- Standard library untuk HTTP client dan JSON

## 🔄 Comparison dengan TypeScript Version

| Feature | TypeScript | Go |
|---------|-----------|-----|
| Performance | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Memory Usage | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| Type Safety | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Concurrency | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Deployment | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

## 🐛 Troubleshooting

**Error: "No MCP servers available"**
- Pastikan MCP server berjalan di URL yang benar
- Check health endpoint: `curl http://localhost:3001/health`

**Error: "Missing api key"**
- Pastikan API key sudah di-set untuk provider yang dipilih
- Verifikasi format API key sudah benar

**Streaming tidak berfungsi:**
- Pastikan client mendukung Server-Sent Events
- Check network/proxy settings

## 📄 License

MIT

## 👥 Contributors

- Jefri Herdi Triyanto (@jefripunza)

## 🔗 Related Projects

- [TypeScript Version](./index.ts)
- [MCP Server](../mcp-server)
