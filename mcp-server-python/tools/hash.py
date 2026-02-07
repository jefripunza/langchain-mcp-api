import hashlib
import hmac
import uuid
try:
    import uuid7
    UUID7_AVAILABLE = True
except ImportError:
    UUID7_AVAILABLE = False

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
    elif version == 5:
        # UUID v5 requires namespace and name
        name = args.get("name", "")
        namespace_str = args.get("namespace", "dns")
        
        # Map namespace string to UUID namespace
        namespace_map = {
            "dns": uuid.NAMESPACE_DNS,
            "url": uuid.NAMESPACE_URL,
            "oid": uuid.NAMESPACE_OID,
            "x500": uuid.NAMESPACE_X500
        }
        namespace = namespace_map.get(namespace_str.lower(), uuid.NAMESPACE_DNS)
        
        if not name:
            return {"error": "UUID v5 requires 'name' parameter"}
        
        generated = str(uuid.uuid5(namespace, name))
    elif version == 7:
        # UUID v7 is time-based (RFC 9562)
        if UUID7_AVAILABLE:
            generated = str(uuid7.create())
        else:
            return {"error": "UUID v7 requires 'uuid7-standard' library. Install with: pip install uuid7-standard"}
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
                    "description": "UUID version (1, 4, 5, atau 7, default: 4)",
                    "enum": [1, 4, 5, 7]
                },
                "name": {
                    "type": "string",
                    "description": "Name untuk UUID v5 (REQUIRED untuk version 5)"
                },
                "namespace": {
                    "type": "string",
                    "description": "Namespace untuk UUID v5 (dns, url, oid, x500, default: dns)",
                    "enum": ["dns", "url", "oid", "x500"]
                }
            }
        },
        "handler": generate_uuid
    }
]
