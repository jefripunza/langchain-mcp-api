# 🐍 MCP Server Python

Model Context Protocol (MCP) Server implementation using Python FastAPI with unique utility tools.

## 🚀 Features

### 🔐 Encoding Tools (8 tools)
- **base64_encode** - Encode text to Base64
- **base64_decode** - Decode Base64 to text
- **url_encode** - URL percent encoding
- **url_decode** - Decode URL encoded text
- **html_encode** - Encode HTML special characters
- **html_decode** - Decode HTML entities
- **json_format** - Format JSON with indentation
- **json_minify** - Minify JSON (remove whitespace)

### 🔒 Hash Tools (6 tools)
- **md5_hash** - Generate MD5 hash
- **sha1_hash** - Generate SHA1 hash
- **sha256_hash** - Generate SHA256 hash
- **sha512_hash** - Generate SHA512 hash
- **hmac_sha256** - Generate HMAC-SHA256 with secret key
- **generate_uuid** - Generate UUID v1 or v4

### 🌐 Network Tools (7 tools)
- **validate_ip** - Validate and analyze IP address (IPv4/IPv6)
- **ip_to_int** - Convert IP address to integer
- **int_to_ip** - Convert integer to IP address
- **parse_url** - Parse URL into components
- **build_url** - Build URL from components
- **dns_lookup** - DNS resolution (hostname to IP)
- **reverse_dns** - Reverse DNS lookup (IP to hostname)

### 🔢 Math Tools (10 tools)
- **calculate_percentage** - Calculate percentage
- **calculate_discount** - Calculate price after discount
- **calculate_compound_interest** - Calculate compound interest
- **calculate_average** - Calculate mean/average
- **calculate_median** - Calculate median
- **calculate_standard_deviation** - Calculate standard deviation
- **calculate_factorial** - Calculate factorial (n!)
- **calculate_gcd** - Calculate GCD (Greatest Common Divisor)
- **calculate_lcm** - Calculate LCM (Least Common Multiple)
- **is_prime** - Check if number is prime

### 📁 File Tools (6 tools)
- **get_file_extension** - Get file extension from filename
- **get_mime_type** - Get MIME type from filename
- **parse_path** - Parse path into components
- **join_path** - Join path parts
- **normalize_path** - Normalize path (resolve '..' and redundant separators)
- **format_bytes** - Format bytes to readable size (KB, MB, GB, etc)

**Total: 37 unique tools**

## 📦 Installation

### ⚠️ Python Version Requirement

**Important:** This project requires **Python 3.9 - 3.13** due to Pydantic dependencies. Python 3.14+ is not yet supported.

Check your Python version:
```bash
python3 --version
```

If you have Python 3.14, install Python 3.13:
```bash
# macOS
brew install python@3.13

# Or use pyenv
pyenv install 3.13.0
pyenv local 3.13.0
```

### Quick Setup (Recommended)

```bash
# Clean setup with compatible Python version
./setup.sh
```

Or use the activation script:
```bash
# Run the activation script (creates venv, installs deps, activates)
./activate.sh
```

### Manual Setup

```bash
# Create virtual environment
python3 -m venv venv

# Activate virtual environment
# On macOS/Linux:
source venv/bin/activate
# On Windows:
# venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt
```

### Using Docker

```bash
# Build image
docker build -t mcp-server-python .

# Run container
docker run -d -p 4050:4050 --name mcp-server-python mcp-server-python
```

## 🏃 Running the Server

### Quick Start (Recommended)

```bash
# Run the server (auto-activates venv)
./run.sh
```

### Manual Start

```bash
# Make sure virtual environment is activated
source venv/bin/activate

# Run the server
python main.py
```

Or with uvicorn directly:

```bash
uvicorn main:app --host 0.0.0.0 --port 4050 --reload
```

### Production Mode

```bash
uvicorn main:app --host 0.0.0.0 --port 4050 --workers 4
```

Server will run on: **http://localhost:4050**

## 📡 API Endpoints

### Health Check

```bash
curl http://localhost:4050/health
```

Response:
```json
{
  "status": "ok"
}
```

### List All Tools

```bash
curl http://localhost:4050/mcp/tools
```

### Invoke Tool

```bash
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "base64_encode",
    "arguments": {
      "text": "Hello World"
    }
  }'
```

## 🧪 Example Usage

### Encoding Example

```bash
# Base64 encode
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "base64_encode",
    "arguments": {"text": "Hello World"}
  }'

# URL encode
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "url_encode",
    "arguments": {"text": "hello world & test"}
  }'
```

### Hash Example

```bash
# SHA256 hash
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "sha256_hash",
    "arguments": {"text": "password123"}
  }'

# Generate UUID
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "generate_uuid",
    "arguments": {"version": 4}
  }'
```

### Network Example

```bash
# Validate IP
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "validate_ip",
    "arguments": {"ip": "192.168.1.1"}
  }'

# DNS lookup
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "dns_lookup",
    "arguments": {"hostname": "google.com"}
  }'
```

### Math Example

```bash
# Calculate percentage
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "calculate_percentage",
    "arguments": {"value": 25, "total": 100}
  }'

# Check prime number
curl -X POST http://localhost:4050/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "is_prime",
    "arguments": {"number": 17}
  }'
```

## 🔗 Integration with LangChain MCP API

Add this server to your LangChain MCP API request:

```bash
curl -X POST http://localhost:6000/chat \
  -H "Content-Type: application/json" \
  -d '{
    "credential": {
      "provider": "openai",
      "api_key": "sk-...",
      "model": "gpt-4o-mini"
    },
    "input": "Please encode 'Hello World' to base64",
    "servers": ["http://localhost:4050"]
  }'
```

## 📝 Project Structure

```
mcp-server-python/
├── main.py                 # FastAPI application entry point
├── requirements.txt        # Python dependencies
├── README.md              # This file
├── .gitignore            # Git ignore rules
└── tools/                # Tool implementations
    ├── encoding.py       # Encoding/decoding tools
    ├── hash.py          # Hashing and UUID tools
    ├── network.py       # Network utility tools
    ├── math_tools.py    # Mathematical calculations
    └── file_tools.py    # File path utilities
```

## 🛠️ Development

### Adding New Tools

1. Create a new file in `tools/` directory (e.g., `tools/my_tools.py`)
2. Define your tool functions and tool definitions
3. Import and register in `main.py`:

```python
from tools.my_tools import my_tools

all_tools.extend(my_tools)
```

### Tool Definition Format

```python
def my_function(args):
    param = args.get("param", "default")
    # Your logic here
    return {"result": "value"}

my_tools = [
    {
        "name": "my_tool",
        "description": "Tool description",
        "parameters": {
            "type": "object",
            "properties": {
                "param": {
                    "type": "string",
                    "description": "Parameter description"
                }
            },
            "required": ["param"]
        },
        "handler": my_function
    }
]
```

## 📄 License

MIT License

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
