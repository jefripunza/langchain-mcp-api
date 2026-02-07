import os
import mimetypes

def get_file_extension(args):
    filename = args.get("filename", "")
    _, ext = os.path.splitext(filename)
    print(f"✅ MCP3 get_file_extension: {filename} -> .{ext.lstrip('.')}")
    return {"extension": ext.lstrip('.')}

def get_mime_type(args):
    filename = args.get("filename", "")
    mime_type, _ = mimetypes.guess_type(filename)
    print(f"✅ MCP3 get_mime_type: {filename} -> {mime_type or 'unknown'}")
    return {"mime_type": mime_type or "unknown"}

def parse_path(args):
    path = args.get("path", "")
    basename = os.path.basename(path)
    print(f"✅ MCP3 parse_path: {path} -> {basename}")
    return {
        "dirname": os.path.dirname(path),
        "basename": basename,
        "filename": os.path.splitext(basename)[0],
        "extension": os.path.splitext(path)[1].lstrip('.'),
        "is_absolute": os.path.isabs(path)
    }

def join_path(args):
    parts = args.get("parts", [])
    if not parts:
        print(f"❌ MCP3 join_path: Parts array cannot be empty")
        return {"error": "Parts array cannot be empty"}
    joined = os.path.join(*parts)
    print(f"✅ MCP3 join_path: {len(parts)} parts -> {joined}")
    return {"path": joined}

def normalize_path(args):
    path = args.get("path", "")
    normalized = os.path.normpath(path)
    print(f"✅ MCP3 normalize_path: {path} -> {normalized}")
    return {"normalized": normalized}

def format_bytes(args):
    bytes_value = args.get("bytes", 0)
    units = args.get("units", "auto")
    original_bytes = bytes_value
    
    if units == "auto":
        for unit in ['B', 'KB', 'MB', 'GB', 'TB', 'PB']:
            if bytes_value < 1024.0:
                formatted = f"{bytes_value:.2f} {unit}"
                print(f"✅ MCP3 format_bytes: {original_bytes} bytes -> {formatted}")
                return {"formatted": formatted, "value": bytes_value, "unit": unit}
            bytes_value /= 1024.0
        formatted = f"{bytes_value:.2f} PB"
        print(f"✅ MCP3 format_bytes: {original_bytes} bytes -> {formatted}")
        return {"formatted": formatted, "value": bytes_value, "unit": "PB"}
    else:
        unit_map = {
            "B": 1,
            "KB": 1024,
            "MB": 1024**2,
            "GB": 1024**3,
            "TB": 1024**4,
            "PB": 1024**5
        }
        if units.upper() in unit_map:
            value = bytes_value / unit_map[units.upper()]
            formatted = f"{value:.2f} {units.upper()}"
            print(f"✅ MCP3 format_bytes: {original_bytes} bytes -> {formatted}")
            return {"formatted": formatted, "value": value, "unit": units.upper()}
        print(f"❌ MCP3 format_bytes: Invalid unit {units}")
        return {"error": "Invalid unit"}

file_tools = [
    {
        "name": "get_file_extension",
        "description": "Dapatkan ekstensi file dari nama file",
        "parameters": {
            "type": "object",
            "properties": {
                "filename": {
                    "type": "string",
                    "description": "Nama file"
                }
            },
            "required": ["filename"]
        },
        "handler": get_file_extension
    },
    {
        "name": "get_mime_type",
        "description": "Dapatkan MIME type dari nama file",
        "parameters": {
            "type": "object",
            "properties": {
                "filename": {
                    "type": "string",
                    "description": "Nama file"
                }
            },
            "required": ["filename"]
        },
        "handler": get_mime_type
    },
    {
        "name": "parse_path",
        "description": "Parse path menjadi komponen-komponennya",
        "parameters": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string",
                    "description": "Path yang akan di-parse"
                }
            },
            "required": ["path"]
        },
        "handler": parse_path
    },
    {
        "name": "join_path",
        "description": "Join beberapa bagian path menjadi satu path",
        "parameters": {
            "type": "object",
            "properties": {
                "parts": {
                    "type": "array",
                    "items": {"type": "string"},
                    "description": "Array bagian path yang akan di-join"
                }
            },
            "required": ["parts"]
        },
        "handler": join_path
    },
    {
        "name": "normalize_path",
        "description": "Normalize path (hapus redundant separator dan resolve '..')",
        "parameters": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string",
                    "description": "Path yang akan di-normalize"
                }
            },
            "required": ["path"]
        },
        "handler": normalize_path
    },
    {
        "name": "format_bytes",
        "description": "Format ukuran bytes ke format yang lebih readable (KB, MB, GB, dll)",
        "parameters": {
            "type": "object",
            "properties": {
                "bytes": {
                    "type": "number",
                    "description": "Ukuran dalam bytes"
                },
                "units": {
                    "type": "string",
                    "description": "Unit target (B, KB, MB, GB, TB, PB) atau 'auto' (default: auto)",
                    "enum": ["auto", "B", "KB", "MB", "GB", "TB", "PB"]
                }
            },
            "required": ["bytes"]
        },
        "handler": format_bytes
    }
]
