import socket
import ipaddress
import urllib.parse

def validate_ip(args):
    ip = args.get("ip", "")
    try:
        ip_obj = ipaddress.ip_address(ip)
        print(f"✅ MCP3 validate_ip: {ip} -> IPv{ip_obj.version}, private={ip_obj.is_private}")
        return {
            "valid": True,
            "version": ip_obj.version,
            "is_private": ip_obj.is_private,
            "is_loopback": ip_obj.is_loopback,
            "is_multicast": ip_obj.is_multicast
        }
    except ValueError:
        print(f"❌ MCP3 validate_ip: {ip} -> Invalid IP address")
        return {"valid": False, "error": "Invalid IP address"}

def ip_to_int(args):
    ip = args.get("ip", "")
    try:
        ip_obj = ipaddress.ip_address(ip)
        result = int(ip_obj)
        print(f"✅ MCP3 ip_to_int: {ip} -> {result}")
        return {"integer": result}
    except ValueError:
        print(f"❌ MCP3 ip_to_int: {ip} -> Invalid IP address")
        return {"error": "Invalid IP address"}

def int_to_ip(args):
    number = args.get("number", 0)
    version = args.get("version", 4)
    try:
        if version == 4:
            ip = ipaddress.IPv4Address(number)
        else:
            ip = ipaddress.IPv6Address(number)
        print(f"✅ MCP3 int_to_ip: {number} -> {str(ip)} (IPv{version})")
        return {"ip": str(ip)}
    except Exception as e:
        print(f"❌ MCP3 int_to_ip: {number} -> {str(e)}")
        return {"error": str(e)}

def parse_url(args):
    url = args.get("url", "")
    try:
        parsed = urllib.parse.urlparse(url)
        print(f"✅ MCP3 parse_url: {url} -> {parsed.scheme}://{parsed.netloc}{parsed.path}")
        return {
            "scheme": parsed.scheme,
            "netloc": parsed.netloc,
            "hostname": parsed.hostname,
            "port": parsed.port,
            "path": parsed.path,
            "params": parsed.params,
            "query": parsed.query,
            "fragment": parsed.fragment
        }
    except Exception as e:
        print(f"❌ MCP3 parse_url: {url} -> {str(e)}")
        return {"error": str(e)}

def build_url(args):
    scheme = args.get("scheme", "https")
    hostname = args.get("hostname", "")
    port = args.get("port")
    path = args.get("path", "")
    query = args.get("query", "")
    fragment = args.get("fragment", "")
    
    netloc = hostname
    if port:
        netloc = f"{hostname}:{port}"
    
    url = urllib.parse.urlunparse((
        scheme,
        netloc,
        path,
        "",
        query,
        fragment
    ))
    print(f"✅ MCP3 build_url: {scheme}://{netloc}{path} -> {url}")
    return {"url": url}

def dns_lookup(args):
    hostname = args.get("hostname", "")
    try:
        ip = socket.gethostbyname(hostname)
        print(f"✅ MCP3 dns_lookup: {hostname} -> {ip}")
        return {"ip": ip, "hostname": hostname}
    except Exception as e:
        print(f"❌ MCP3 dns_lookup: {hostname} -> {str(e)}")
        return {"error": str(e)}

def reverse_dns(args):
    ip = args.get("ip", "")
    try:
        hostname = socket.gethostbyaddr(ip)[0]
        print(f"✅ MCP3 reverse_dns: {ip} -> {hostname}")
        return {"hostname": hostname, "ip": ip}
    except socket.herror as e:
        print(f"❌ MCP3 reverse_dns: {ip} -> {str(e)}")
        return {"error": str(e)}

network_tools = [
    {
        "name": "validate_ip",
        "description": "Validasi dan analisis IP address (IPv4/IPv6)",
        "parameters": {
            "type": "object",
            "properties": {
                "ip": {
                    "type": "string",
                    "description": "IP address yang akan divalidasi"
                }
            },
            "required": ["ip"]
        },
        "handler": validate_ip
    },
    {
        "name": "ip_to_int",
        "description": "Convert IP address ke integer",
        "parameters": {
            "type": "object",
            "properties": {
                "ip": {
                    "type": "string",
                    "description": "IP address yang akan dikonversi"
                }
            },
            "required": ["ip"]
        },
        "handler": ip_to_int
    },
    {
        "name": "int_to_ip",
        "description": "Convert integer ke IP address",
        "parameters": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "integer",
                    "description": "Integer yang akan dikonversi"
                },
                "version": {
                    "type": "integer",
                    "description": "IP version (4 atau 6, default: 4)",
                    "enum": [4, 6]
                }
            },
            "required": ["number"]
        },
        "handler": int_to_ip
    },
    {
        "name": "parse_url",
        "description": "Parse URL menjadi komponen-komponennya",
        "parameters": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string",
                    "description": "URL yang akan di-parse"
                }
            },
            "required": ["url"]
        },
        "handler": parse_url
    },
    {
        "name": "build_url",
        "description": "Build URL dari komponen-komponennya",
        "parameters": {
            "type": "object",
            "properties": {
                "scheme": {
                    "type": "string",
                    "description": "URL scheme (http, https, dll)"
                },
                "hostname": {
                    "type": "string",
                    "description": "Hostname atau domain"
                },
                "port": {
                    "type": "integer",
                    "description": "Port number (optional)"
                },
                "path": {
                    "type": "string",
                    "description": "URL path (optional)"
                },
                "query": {
                    "type": "string",
                    "description": "Query string (optional)"
                },
                "fragment": {
                    "type": "string",
                    "description": "URL fragment (optional)"
                }
            },
            "required": ["hostname"]
        },
        "handler": build_url
    },
    {
        "name": "dns_lookup",
        "description": "Lookup IP address dari hostname (DNS resolution)",
        "parameters": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string",
                    "description": "Hostname yang akan di-lookup"
                }
            },
            "required": ["hostname"]
        },
        "handler": dns_lookup
    },
    {
        "name": "reverse_dns",
        "description": "Reverse DNS lookup (IP ke hostname)",
        "parameters": {
            "type": "object",
            "properties": {
                "ip": {
                    "type": "string",
                    "description": "IP address yang akan di-lookup"
                }
            },
            "required": ["ip"]
        },
        "handler": reverse_dns
    }
]
