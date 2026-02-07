# MCP Server - PHP Laravel 11

MCP (Model Context Protocol) Server menggunakan Laravel 11 framework dengan FrankenPHP.

## Setup

### Prerequisites
- PHP 8.2+
- Composer
- FrankenPHP

### Installation

1. Install dependencies:
```bash
composer install
```

2. Copy environment file:
```bash
cp .env.example .env
```

3. Generate application key:
```bash
php artisan key:generate
```

4. Run server:
```bash
frankenphp run
```

Server akan berjalan di http://localhost:4090

## API Endpoints

### 1. Root Endpoint
GET /
```bash
curl http://localhost:4090/
```
Response:
```json
{
  "message": "MCP Server PHP Laravel is running"
}
```

### 2. Health Check
GET /health
```bash
curl http://localhost:4090/health
```
Response:
```json
{
  "status": "ok"
}
```

### 3. List Tools
GET /mcp/tools
```bash
curl http://localhost:4090/mcp/tools
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
POST /mcp/invoke
```bash
curl -X POST http://localhost:4090/mcp/invoke \
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

## Available Tools

### String Tools (5 tools)
- string_reverse - Membalikkan urutan karakter
- string_uppercase - Ubah ke huruf besar
- string_lowercase - Ubah ke huruf kecil
- string_length - Hitung panjang string
- string_trim - Hapus whitespace

### Math Tools (6 tools)
- math_add - Penjumlahan (a + b)
- math_subtract - Pengurangan (a - b)
- math_multiply - Perkalian (a x b)
- math_divide - Pembagian (a / b)
- math_power - Pangkat (a^b)
- math_sqrt - Akar kuadrat

### Network Tools (3 tools)
- network_validate_ip - Validasi IP address
- network_dns_lookup - DNS lookup hostname ke IP
- network_parse_url - Parse URL ke komponen

### File Tools (3 tools)
- file_get_extension - Dapatkan ekstensi file
- file_get_basename - Dapatkan nama file tanpa path
- file_format_bytes - Format bytes ke KB/MB/GB

## Example Usage

### String Reverse
```bash
curl -X POST http://localhost:4090/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "string_reverse",
    "arguments": {"text": "Laravel"}
  }'
```

### Math Add
```bash
curl -X POST http://localhost:4090/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "math_add",
    "arguments": {"a": 10, "b": 20}
  }'
```

### DNS Lookup
```bash
curl -X POST http://localhost:4090/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "network_dns_lookup",
    "arguments": {"hostname": "google.com"}
  }'
```

### Format Bytes
```bash
curl -X POST http://localhost:4090/mcp/invoke \
  -H "Content-Type: application/json" \
  -d '{
    "name": "file_format_bytes",
    "arguments": {"bytes": 1048576}
  }'
```

## Project Structure

```
app/
├── Http/
│   └── Controllers/
│       └── McpController.php       # Main MCP API controller
└── Services/
    ├── ToolRegistry.php            # Tools registry manager
    └── Tools/
        ├── StringTools.php         # String manipulation tools
        ├── MathTools.php           # Mathematical operations
        ├── NetworkTools.php        # Network utilities
        └── FileTools.php           # File utilities
routes/
└── web.php                         # API routes configuration
```

## Adding New Tools

1. Create new tool class in app/Services/Tools/:
```php
<?php
namespace App\Services\Tools;

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
        return ['result' => 'value'];
    }
}
```

2. Register in app/Services/ToolRegistry.php:
```php
use App\Services\Tools\YourTools;

// In registerTools() method:
$yourTools = new YourTools();
foreach ($yourTools->getTools() as $tool) {
    $this->tools[] = $tool;
}
```

## Integration dengan LangChain Server

Tambahkan server ini ke LangChain MCP API:
```json
{
  "mcp_servers": ["http://localhost:4090"]
}
```

## License

MIT License
