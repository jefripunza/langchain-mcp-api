import base64
import urllib.parse
import html
import json

def base64_encode(args):
    text = args.get("text", "")
    encoded = base64.b64encode(text.encode()).decode()
    print(f"✅ MCP3 base64_encode: {len(text)} chars -> {len(encoded)} chars")
    return {"encoded": encoded}

def base64_decode(args):
    text = args.get("text", "")
    try:
        decoded = base64.b64decode(text).decode()
        print(f"✅ MCP3 base64_decode: {len(text)} chars -> {len(decoded)} chars")
        return {"decoded": decoded}
    except Exception as e:
        print(f"❌ MCP3 base64_decode: {str(e)}")
        return {"error": str(e)}

def url_encode(args):
    text = args.get("text", "")
    encoded = urllib.parse.quote(text)
    print(f"✅ MCP3 url_encode: {text[:50]}... -> {encoded[:50]}...")
    return {"encoded": encoded}

def url_decode(args):
    text = args.get("text", "")
    decoded = urllib.parse.unquote(text)
    print(f"✅ MCP3 url_decode: {text[:50]}... -> {decoded[:50]}...")
    return {"decoded": decoded}

def html_encode(args):
    text = args.get("text", "")
    encoded = html.escape(text)
    print(f"✅ MCP3 html_encode: {len(text)} chars")
    return {"encoded": encoded}

def html_decode(args):
    text = args.get("text", "")
    decoded = html.unescape(text)
    print(f"✅ MCP3 html_decode: {len(text)} chars")
    return {"decoded": decoded}

def json_format(args):
    text = args.get("text", "")
    indent = args.get("indent", 2)
    try:
        parsed = json.loads(text)
        formatted = json.dumps(parsed, indent=indent, ensure_ascii=False)
        print(f"✅ MCP3 json_format: indent={indent}")
        return {"formatted": formatted}
    except Exception as e:
        print(f"❌ MCP3 json_format: {str(e)}")
        return {"error": str(e)}

def json_minify(args):
    text = args.get("text", "")
    try:
        parsed = json.loads(text)
        minified = json.dumps(parsed, separators=(',', ':'), ensure_ascii=False)
        print(f"✅ MCP3 json_minify: {len(text)} -> {len(minified)} chars")
        return {"minified": minified}
    except Exception as e:
        print(f"❌ MCP3 json_minify: {str(e)}")
        return {"error": str(e)}

encoding_tools = [
    {
        "name": "base64_encode",
        "description": "Encode text ke Base64",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-encode"
                }
            },
            "required": ["text"]
        },
        "handler": base64_encode
    },
    {
        "name": "base64_decode",
        "description": "Decode Base64 ke text",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Base64 string yang akan di-decode"
                }
            },
            "required": ["text"]
        },
        "handler": base64_decode
    },
    {
        "name": "url_encode",
        "description": "Encode text untuk URL (percent encoding)",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-encode"
                }
            },
            "required": ["text"]
        },
        "handler": url_encode
    },
    {
        "name": "url_decode",
        "description": "Decode URL encoded text",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "URL encoded text yang akan di-decode"
                }
            },
            "required": ["text"]
        },
        "handler": url_decode
    },
    {
        "name": "html_encode",
        "description": "Encode special characters untuk HTML",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-encode"
                }
            },
            "required": ["text"]
        },
        "handler": html_encode
    },
    {
        "name": "html_decode",
        "description": "Decode HTML entities ke text",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "HTML encoded text yang akan di-decode"
                }
            },
            "required": ["text"]
        },
        "handler": html_decode
    },
    {
        "name": "json_format",
        "description": "Format JSON string dengan indentasi",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "JSON string yang akan di-format"
                },
                "indent": {
                    "type": "integer",
                    "description": "Jumlah spasi untuk indentasi (default: 2)"
                }
            },
            "required": ["text"]
        },
        "handler": json_format
    },
    {
        "name": "json_minify",
        "description": "Minify JSON string (hapus whitespace)",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "JSON string yang akan di-minify"
                }
            },
            "required": ["text"]
        },
        "handler": json_minify
    }
]
