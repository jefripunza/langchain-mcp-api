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

## 🐳 Docker

### Option 1: Pull from Docker Hub (Recommended)

Pull the pre-built image from Docker Hub:

```bash
# Pull the latest image
docker pull jefriherditriyanto/langchain-mcp-api:latest

# Run container
docker run -d \
  --name langchain-mcp-api \
  -p 6000:6000 \
  jefriherditriyanto/langchain-mcp-api:latest
```

**Using Docker Compose:**

```yaml
version: '3.8'
services:
  langchain-mcp-api:
    image: jefriherditriyanto/langchain-mcp-api:latest
    container_name: langchain-mcp-api
    ports:
      - "6000:6000"
    environment:
      - PORT=6000
    restart: unless-stopped
```

Run with:
```bash
docker-compose up -d
```

---

## 📡 API Documentation

### Base URL
```
http://localhost:6000
```

### Endpoints

#### 1️⃣ **Hello World**

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

#### 2️⃣ **Health Check**

```http
GET /health
```

**Response:**
```json
{
  "status": "ok"
}
```

---

#### 3️⃣ **Chat (Non-Streaming)**

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
  "servers": ["http://host.docker.internal:4000"]
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

#### 4️⃣ **Chat Stream (SSE)**

```http
POST /chat/stream
```

**Request Body:** *(same as `/chat`)*

**Response:** Server-Sent Events stream

```
data: {"type":"start","timestamp":"2024-02-04T09:00:00Z","input":"What is the weather?"}

data: {"type":"servers_checked","available_servers":["http://host.docker.internal:4000"],"total_servers":1}

data: {"type":"thinking_start","timestamp":"2024-02-04T09:00:01Z"}

data: {"type":"thinking_chunk","chunk":"I need to check the weather...","is_final":false}

data: {"type":"message_start","timestamp":"2024-02-04T09:00:02Z"}

data: {"type":"message_chunk","chunk":"The weather is ","is_final":false}

data: {"type":"message_chunk","chunk":"sunny, 28°C","is_final":true}

data: {"type":"done","done":true,"total_steps":3,"timestamp":"2024-02-04T09:00:03Z"}
```

---

## ⚙️ Configuration

### Provider Settings

| Provider       | Key          | Required Fields    |
| -------------- | ------------ | ------------------ |
| **OpenAI**     | `openai`     | `api_key`, `model` |
| **Claude**     | `claude`     | `api_key`, `model` |
| **OpenRouter** | `openrouter` | `api_key`, `model` |
| **Ollama**     | `ollama`     | `url`, `model`     |
| **Llama.cpp**  | `llama_cpp`  | `url`, `model`     |
| **vLLM**       | `vllm`       | `url`, `model`     |

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

Control detailed execution logs with the `VERBOSE` environment variable:

### Enable Verbose Mode

```bash
# Enable verbose logging (shows all requestID-prefixed logs)
export VERBOSE=true
go run main.go

# Or with Docker
docker run -d \
  -p 6000:6000 \
  -e VERBOSE=true \
  jefriherditriyanto/langchain-mcp-api:latest
```

### Disable Verbose Mode

```bash
# Disable verbose logging (hides requestID-prefixed logs)
export VERBOSE=false
go run main.go

# Or simply don't set the variable (defaults to false)
go run main.go
```

### Example Verbose Output

When `VERBOSE=true`, you'll see detailed execution traces:

```
[1kqlh2PxHZZvLVLiBbdbWfdXQ9] [START REQUEST]
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]📦 [AGENT] Creating LangChain Agent...
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]   Provider: llama_cpp
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]   Model: gpt-oss-20b.gguf
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]   ✅ Loaded 22 tools from MCP servers

[1kqlh2PxHZZvLVLiBbdbWfdXQ9]🚀 [INVOKE] Starting agent invocation...
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]   Input: What is the weather?

[1kqlh2PxHZZvLVLiBbdbWfdXQ9]   🔁 [ITERATION 1/10]
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]      📝 Built 2 messages for LLM
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]      🤖 Calling LLM...
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]      ✅ LLM Response (245 chars)
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]      🔧 Detected 1 tool call(s)
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]         1. getWeather({"lat": -7.7, "lon": 109.0})
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]      ⚙️  Executing tools...
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]         [1/1] Executing: getWeather
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]            ✅ Success from http://host.docker.internal:4000
[1kqlh2PxHZZvLVLiBbdbWfdXQ9]      ✅ Tools executed successfully (1 results)

[1kqlh2PxHZZvLVLiBbdbWfdXQ9] ✅ [INVOKE] Agent invocation completed
```

---

## � MCP Server Example

Build your own MCP (Model Context Protocol) server to provide custom tools for the LangChain API.

### Complete Example

Full working example available at: [mcp-server-typescript](https://github.com/jefripunza/langchain-mcp-api/tree/master/mcp-server-typescript) [mcp-server-golang](https://github.com/jefripunza/langchain-mcp-api/tree/master/mcp-server-golang)

### Quick Start

```bash
# Clone MCP server example
git clone https://github.com/jefripunza/langchain-mcp-api/tree/master/mcp-server-typescript

# Navigate to MCP server example
cd mcp-server-typescript

# Install dependencies
bun install

# Run the server
bun run dev
```

Server will start at `http://localhost:4000` 🎉

---

### Project Structure

```
mcp-server-typescript/
├── src/
│   ├── index.ts         # Main server
│   ├── registry.ts      # Tool registry
│   └── tools/
│       ├── math.ts      # Math tools
│       └── weather.ts   # Weather tools
├── package.json
└── tsconfig.json
```

---

### Implementation Guide

#### 1️⃣ **Main Server** (`src/index.ts`)

```typescript
import express from "express";
import cors from "cors";
import helmet from "helmet";
import morgan from "morgan";

import { tools, findTool } from "./registry";

const app = express();
app.use(express.json());
app.use(cors());
app.use(helmet());

app.listen(4000, () => {
  console.log("🧠 MCP Server running on http://localhost:4000");
});
app.use(morgan("dev"));

// REQUIRED!
app.get("/health", (_req, res) => res.json({ status: "ok" }));

// MCP-style: list tools, REQUIRED!
app.get("/mcp/tools", (_req, res) => {
  res.json(
    tools.map((t) => ({
      name: t.name,
      description: t.description,
      parameters: t.parameters,
    })),
  );
});

// MCP-style: invoke tool, REQUIRED!
app.post("/mcp/invoke", async (req, res) => {
  const { name, arguments: args } = req.body;
  const tool = findTool(name);

  if (!tool) {
    return res.status(404).json({ error: "Tool not found" });
  }
  if (!tool.handler) {
    return res.status(400).json({ error: "Tool handler not found" });
  }

  const result = await tool.handler(args);
  res.json(result);
});
```

---

#### 2️⃣ **Tool Registry** (`src/registry.ts`)

```typescript
import type { Tool } from "../../types/tool";
import { mathTools } from "./tools/math";
import { weatherTools } from "./tools/weather";

export const tools: Tool[] = [...mathTools, ...weatherTools];

export function findTool(name: string) {
  return tools.find((t) => t.name === name);
}
```

---

#### 3️⃣ **Math Tool Example** (`src/tools/math.ts`)

```typescript
import type { Tool } from "../../../types/tool";

export const mathTools: Tool[] = [
  {
    name: "add",
    description: "Add two numbers together", // add to main prompt, please detail!
    // add to main prompt, please detail!
    parameters: {
      type: "object",
      properties: {
        a: { type: "number" },
        b: { type: "number" },
      },
      required: ["a", "b"],
    },
    // as controller / logic base
    handler: async ({ a, b }: { a: number; b: number }) => {
      console.log(`✅ MCP1 Math: ${a}+${b}=${a + b}`);
      return { result: a + b };
    },
  },
];
```

---

#### 4️⃣ **Weather Tool Example** (`src/tools/weather.ts`)

```typescript
import { fetchWeatherApi } from "openmeteo";
import type { Tool } from "../../../types/tool";

// advance tool version
export const weatherTools: Tool[] = [
  {
    name: "getWeather",
    description: "Get weather data by coordinates",
    parameters: {
      type: "object",
      properties: {
        latitude: { type: "number" },
        longitude: { type: "number" },
      },
      required: ["latitude", "longitude"],
    },
    handler: async ({
      latitude,
      longitude,
    }: {
      latitude: number;
      longitude: number;
    }) => {
      const params = {
        latitude,
        longitude,
        hourly: ["temperature_2m", "relative_humidity_2m", "rain"],
        timezone: "auto",
      };
      
      const responses = await fetchWeatherApi(
        "https://api.open-meteo.com/v1/forecast",
        params
      );
      
      const response = responses[0];
      const hourly = response.hourly()!;
      
      console.log(`✅ MCP1 Weather: ${latitude}, ${longitude}`);
      return {
        latitude,
        longitude,
        temperature: hourly.variables(0)!.valuesArray(),
        humidity: hourly.variables(1)!.valuesArray(),
        rain: hourly.variables(2)!.valuesArray(),
      };
    },
  },
];
```

---

### MCP Protocol Endpoints

Your MCP server must implement these three endpoints very required:

| Endpoint      | Method | Description              |
| ------------- | ------ | ------------------------ |
| `/health`     | GET    | Health check             |
| `/mcp/tools`  | GET    | List all available tools |
| `/mcp/invoke` | POST   | Execute a specific tool  |

---

### Testing Your MCP Server

```bash
# List available tools
curl http://localhost:4000/mcp/tools

# Invoke math tool
curl -X POST http://localhost:4000/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "add",
    "arguments": {"a": 5, "b": 3}
  }'

# Invoke weather tool
curl -X POST http://localhost:4000/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "getWeather",
    "arguments": {"latitude": -6.2, "longitude": 106.8}
  }'
```

---

### Using with LangChain API

Once your MCP server is running, use it with the LangChain API:

```bash
curl -X POST http://localhost:6000/chat \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "openai",
      "api_key": "sk-...",
      "model": "gpt-4o-mini"
    },
    "input": "What is 25 + 37?",
    "servers": ["http://host.docker.internal:4000"]
  }'
```

The LangChain API will automatically:
1. Discover tools from your MCP server
2. Let the LLM decide which tools to use
3. Execute the tools and return results

---

## 🐛 Troubleshooting

<details>
<summary><b>Error: "No MCP servers available"</b></summary>

- Ensure MCP server is running
- Check server URL is correct
- Verify health endpoint: `curl http://localhost:4000/health`
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

| Metric                  | Value                 |
| ----------------------- | --------------------- |
| **Startup Time**        | < 1s                  |
| **Memory Usage**        | ~50MB (idle)          |
| **Concurrent Requests** | 1000+                 |
| **Response Time**       | < 100ms (without LLM) |

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
