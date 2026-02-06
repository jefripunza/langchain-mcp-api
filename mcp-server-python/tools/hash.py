import hashlib
import hmac
import uuid

def md5_hash(args):
    text = args.get("text", "")
    hashed = hashlib.md5(text.encode()).hexdigest()
    return {"hash": hashed}

def sha1_hash(args):
    text = args.get("text", "")
    hashed = hashlib.sha1(text.encode()).hexdigest()
    return {"hash": hashed}

def sha256_hash(args):
    text = args.get("text", "")
    hashed = hashlib.sha256(text.encode()).hexdigest()
    return {"hash": hashed}

def sha512_hash(args):
    text = args.get("text", "")
    hashed = hashlib.sha512(text.encode()).hexdigest()
    return {"hash": hashed}

def hmac_sha256(args):
    text = args.get("text", "")
    key = args.get("key", "")
    hashed = hmac.new(key.encode(), text.encode(), hashlib.sha256).hexdigest()
    return {"hash": hashed}

def generate_uuid(args):
    version = args.get("version", 4)
    if version == 1:
        generated = str(uuid.uuid1())
    elif version == 4:
        generated = str(uuid.uuid4())
    else:
        generated = str(uuid.uuid4())
    return {"uuid": generated}

hash_tools = [
    {
        "name": "md5_hash",
        "description": "Generate MD5 hash dari text",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-hash"
                }
            },
            "required": ["text"]
        },
        "handler": md5_hash
    },
    {
        "name": "sha1_hash",
        "description": "Generate SHA1 hash dari text",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-hash"
                }
            },
            "required": ["text"]
        },
        "handler": sha1_hash
    },
    {
        "name": "sha256_hash",
        "description": "Generate SHA256 hash dari text",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-hash"
                }
            },
            "required": ["text"]
        },
        "handler": sha256_hash
    },
    {
        "name": "sha512_hash",
        "description": "Generate SHA512 hash dari text",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-hash"
                }
            },
            "required": ["text"]
        },
        "handler": sha512_hash
    },
    {
        "name": "hmac_sha256",
        "description": "Generate HMAC-SHA256 hash dengan secret key",
        "parameters": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "description": "Text yang akan di-hash"
                },
                "key": {
                    "type": "string",
                    "description": "Secret key untuk HMAC"
                }
            },
            "required": ["text", "key"]
        },
        "handler": hmac_sha256
    },
    {
        "name": "generate_uuid",
        "description": "Generate UUID (Universally Unique Identifier)",
        "parameters": {
            "type": "object",
            "properties": {
                "version": {
                    "type": "integer",
                    "description": "UUID version (1 atau 4, default: 4)",
                    "enum": [1, 4]
                }
            }
        },
        "handler": generate_uuid
    }
]
