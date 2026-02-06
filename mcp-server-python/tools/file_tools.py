import os
import mimetypes

def get_file_extension(args):
    filename = args.get("filename", "")
    _, ext = os.path.splitext(filename)
    return {"extension": ext.lstrip('.')}

def get_mime_type(args):
    filename = args.get("filename", "")
    mime_type, _ = mimetypes.guess_type(filename)
    return {"mime_type": mime_type or "unknown"}

def parse_path(args):
    path = args.get("path", "")
    return {
        "dirname": os.path.dirname(path),
        "basename": os.path.basename(path),
        "filename": os.path.splitext(os.path.basename(path))[0],
        "extension": os.path.splitext(path)[1].lstrip('.'),
        "is_absolute": os.path.isabs(path)
    }

def join_path(args):
    parts = args.get("parts", [])
    if not parts:
        return {"error": "Parts array cannot be empty"}
    joined = os.path.join(*parts)
    return {"path": joined}

def normalize_path(args):
    path = args.get("path", "")
    normalized = os.path.normpath(path)
    return {"normalized": normalized}

def format_bytes(args):
    bytes_value = args.get("bytes", 0)
    units = args.get("units", "auto")
    
    if units == "auto":
        for unit in ['B', 'KB', 'MB', 'GB', 'TB', 'PB']:
            if bytes_value < 1024.0:
                return {"formatted": f"{bytes_value:.2f} {unit}", "value": bytes_value, "unit": unit}
            bytes_value /= 1024.0
        return {"formatted": f"{bytes_value:.2f} PB", "value": bytes_value, "unit": "PB"}
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
            return {"formatted": f"{value:.2f} {units.upper()}", "value": value, "unit": units.upper()}
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
