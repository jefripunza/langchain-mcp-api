# MCP Server - PHP CodeIgniter 4

MCP (Model Context Protocol) Server menggunakan CodeIgniter 4 framework dengan FrankenPHP.

## 🚀 Setup

### Prerequisites
- PHP 8.1+
- Composer
- FrankenPHP

### Installation

1. Copy environment file:
```bash
cp env .env
```

2. Install dependencies:
```bash
composer install
```

3. Run server:
```bash
frankenphp run
```

Server akan berjalan di `http://localhost:4080`

## 📡 API Endpoints

### 1. Root Endpoint
**GET /** 
```bash
curl http://localhost:4080/
```
Response:
```json
{
  "message": "🧠 MCP Server PHP CodeIgniter is running"
}
```

### 2. Health Check
**GET /health**
```bash
curl http://localhost:4080/health
```
Response:
```json
{
  "status": "ok"
}
```

### 3. List Tools
**GET /mcp/tools**
```bash
curl http://localhost:4080/mcp/tools
```
Response:
```json
[
  {
    "name": "string_reverse",
    "description": "Membalikkan urutan karakter dalam string",
    "parameters": {
      "type": "object",
      "properties": {
        "text": {
          "type": "string",
          "description": "Text yang akan dibalik"
        }
      },
      "required": ["text"]
    }
  }
]
```

### 4. Invoke Tool
**POST /mcp/invoke**
```bash
curl -X POST http://localhost:4080/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "string_reverse",
    "arguments": {
      "text": "hello"
    }
  }'
```
Response:
```json
{
  "result": {
    "reversed": "olleh"
  }
}
```

## 🛠️ Available Tools

### String Tools (5 tools)
- **string_reverse** - Membalikkan urutan karakter
- **string_uppercase** - Ubah ke huruf besar
- **string_lowercase** - Ubah ke huruf kecil
- **string_length** - Hitung panjang string
- **string_trim** - Hapus whitespace

### Math Tools (6 tools)
- **math_add** - Penjumlahan (a + b)
- **math_subtract** - Pengurangan (a - b)
- **math_multiply** - Perkalian (a × b)
- **math_divide** - Pembagian (a ÷ b)
- **math_power** - Pangkat (a^b)
- **math_sqrt** - Akar kuadrat (√n)

### Network Tools (3 tools)
- **network_validate_ip** - Validasi IP address
- **network_dns_lookup** - DNS lookup hostname ke IP
- **network_parse_url** - Parse URL ke komponen

### File Tools (3 tools)
- **file_get_extension** - Dapatkan ekstensi file
- **file_get_basename** - Dapatkan nama file tanpa path
- **file_format_bytes** - Format bytes ke KB/MB/GB

## 📝 Example Usage

### String Reverse
```bash
curl -X POST http://localhost:4080/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "string_reverse",
    "arguments": {"text": "CodeIgniter"}
  }'
```

### Math Add
```bash
curl -X POST http://localhost:4080/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "math_add",
    "arguments": {"a": 10, "b": 20}
  }'
```

### DNS Lookup
```bash
curl -X POST http://localhost:4080/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "network_dns_lookup",
    "arguments": {"hostname": "google.com"}
  }'
```

### Format Bytes
```bash
curl -X POST http://localhost:4080/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "file_format_bytes",
    "arguments": {"bytes": 1048576}
  }'
```

## 🏗️ Project Structure

```
app/
├── Controllers/
│   └── McpController.php       # Main MCP API controller
├── Libraries/
│   ├── ToolRegistry.php        # Tools registry manager
│   └── Tools/
│       ├── StringTools.php     # String manipulation tools
│       ├── MathTools.php       # Mathematical operations
│       ├── NetworkTools.php    # Network utilities
│       └── FileTools.php       # File utilities
└── Config/
    └── Routes.php              # API routes configuration
```

## 🔧 Adding New Tools

1. Create new tool class in `app/Libraries/Tools/`:
```php
<?php
namespace App\Libraries\Tools;

class YourTools
{
    public function getTools()
    {
        return [
            [
                'name' => 'your_tool',
                'description' => 'Tool description',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'param' => [
                            'type' => 'string',
                            'description' => 'Parameter description'
                        ]
                    ],
                    'required' => ['param']
                ],
                'handler' => [$this, 'yourMethod']
            ]
        ];
    }

    public function yourMethod($args)
    {
        $param = $args['param'] ?? '';
        echo "✅ MCP4 your_tool: {$param}\n";
        return ['result' => 'value'];
    }
}
```

2. Register in `app/Libraries/ToolRegistry.php`:
```php
use App\Libraries\Tools\YourTools;

// In registerTools() method:
$yourTools = new YourTools();
foreach ($yourTools->getTools() as $tool) {
    $this->tools[] = $tool;
}
```

## 📊 Logging

Semua tool execution akan menampilkan log dengan format:
```
✅ MCP4 tool_name: input -> output
❌ MCP4 tool_name: error_message
```

## 🔗 Integration dengan LangChain Server

Tambahkan server ini ke LangChain MCP API:
```json
{
  "mcp_servers": ["http://localhost:4080"]
}
```

## 📄 License

MIT License
