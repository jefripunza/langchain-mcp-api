import math
import statistics

def calculate_percentage(args):
    value = args.get("value", 0)
    total = args.get("total", 0)
    if total == 0:
        print(f"❌ MCP3 calculate_percentage: Total cannot be zero")
        return {"error": "Total cannot be zero"}
    percentage = (value / total) * 100
    print(f"✅ MCP3 calculate_percentage: {value}/{total} = {round(percentage, 2)}%")
    return {"percentage": round(percentage, 2)}

def calculate_discount(args):
    original_price = args.get("original_price", 0)
    discount_percent = args.get("discount_percent", 0)
    discount_amount = original_price * (discount_percent / 100)
    final_price = original_price - discount_amount
    print(f"✅ MCP3 calculate_discount: {original_price} - {discount_percent}% = {round(final_price, 2)}")
    return {
        "original_price": original_price,
        "discount_percent": discount_percent,
        "discount_amount": round(discount_amount, 2),
        "final_price": round(final_price, 2)
    }

def calculate_compound_interest(args):
    principal = args.get("principal", 0)
    rate = args.get("rate", 0)
    time = args.get("time", 0)
    frequency = args.get("frequency", 1)
    
    amount = principal * math.pow((1 + rate / (100 * frequency)), frequency * time)
    interest = amount - principal
    print(f"✅ MCP3 calculate_compound_interest: P={principal}, r={rate}%, t={time}y -> {round(amount, 2)}")
    
    return {
        "principal": principal,
        "rate": rate,
        "time": time,
        "frequency": frequency,
        "final_amount": round(amount, 2),
        "interest_earned": round(interest, 2)
    }

def calculate_average(args):
    numbers = args.get("numbers", [])
    if not numbers:
        print(f"❌ MCP3 calculate_average: Numbers array cannot be empty")
        return {"error": "Numbers array cannot be empty"}
    
    avg = statistics.mean(numbers)
    print(f"✅ MCP3 calculate_average: {len(numbers)} numbers -> avg={round(avg, 2)}")
    return {"average": round(avg, 2), "count": len(numbers)}

def calculate_median(args):
    numbers = args.get("numbers", [])
    if not numbers:
        print(f"❌ MCP3 calculate_median: Numbers array cannot be empty")
        return {"error": "Numbers array cannot be empty"}
    
    median = statistics.median(numbers)
    print(f"✅ MCP3 calculate_median: {len(numbers)} numbers -> median={round(median, 2)}")
    return {"median": round(median, 2), "count": len(numbers)}

def calculate_standard_deviation(args):
    numbers = args.get("numbers", [])
    if len(numbers) < 2:
        print(f"❌ MCP3 calculate_standard_deviation: Need at least 2 numbers")
        return {"error": "Need at least 2 numbers"}
    
    std_dev = statistics.stdev(numbers)
    print(f"✅ MCP3 calculate_standard_deviation: {len(numbers)} numbers -> std_dev={round(std_dev, 2)}")
    return {"standard_deviation": round(std_dev, 2), "count": len(numbers)}

def calculate_factorial(args):
    number = args.get("number", 0)
    if number < 0:
        print(f"❌ MCP3 calculate_factorial: Factorial not defined for negative numbers")
        return {"error": "Factorial is not defined for negative numbers"}
    if number > 170:
        print(f"❌ MCP3 calculate_factorial: Number too large (max: 170)")
        return {"error": "Number too large (max: 170)"}
    
    result = math.factorial(number)
    print(f"✅ MCP3 calculate_factorial: {number}! = {result}")
    return {"factorial": result}

def calculate_gcd(args):
    a = args.get("a", 0)
    b = args.get("b", 0)
    result = math.gcd(a, b)
    print(f"✅ MCP3 calculate_gcd: gcd({a}, {b}) = {result}")
    return {"gcd": result, "a": a, "b": b}

def calculate_lcm(args):
    a = args.get("a", 0)
    b = args.get("b", 0)
    result = abs(a * b) // math.gcd(a, b) if a and b else 0
    print(f"✅ MCP3 calculate_lcm: lcm({a}, {b}) = {result}")
    return {"lcm": result, "a": a, "b": b}

def is_prime(args):
    number = args.get("number", 0)
    if number < 2:
        print(f"✅ MCP3 is_prime: {number} -> False (< 2)")
        return {"is_prime": False, "number": number}
    
    for i in range(2, int(math.sqrt(number)) + 1):
        if number % i == 0:
            print(f"✅ MCP3 is_prime: {number} -> False (divisible by {i})")
            return {"is_prime": False, "number": number}
    
    print(f"✅ MCP3 is_prime: {number} -> True")
    return {"is_prime": True, "number": number}

math_tools = [
    {
        "name": "calculate_percentage",
        "description": "Hitung persentase dari nilai terhadap total",
        "parameters": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "number",
                    "description": "Nilai yang akan dihitung persentasenya"
                },
                "total": {
                    "type": "number",
                    "description": "Total nilai"
                }
            },
            "required": ["value", "total"]
        },
        "handler": calculate_percentage
    },
    {
        "name": "calculate_discount",
        "description": "Hitung harga setelah diskon",
        "parameters": {
            "type": "object",
            "properties": {
                "original_price": {
                    "type": "number",
                    "description": "Harga asli"
                },
                "discount_percent": {
                    "type": "number",
                    "description": "Persentase diskon"
                }
            },
            "required": ["original_price", "discount_percent"]
        },
        "handler": calculate_discount
    },
    {
        "name": "calculate_compound_interest",
        "description": "Hitung bunga majemuk (compound interest)",
        "parameters": {
            "type": "object",
            "properties": {
                "principal": {
                    "type": "number",
                    "description": "Modal awal"
                },
                "rate": {
                    "type": "number",
                    "description": "Suku bunga per tahun (%)"
                },
                "time": {
                    "type": "number",
                    "description": "Waktu dalam tahun"
                },
                "frequency": {
                    "type": "integer",
                    "description": "Frekuensi bunga per tahun (default: 1)"
                }
            },
            "required": ["principal", "rate", "time"]
        },
        "handler": calculate_compound_interest
    },
    {
        "name": "calculate_average",
        "description": "Hitung rata-rata (mean) dari array angka",
        "parameters": {
            "type": "object",
            "properties": {
                "numbers": {
                    "type": "array",
                    "items": {"type": "number"},
                    "description": "Array angka"
                }
            },
            "required": ["numbers"]
        },
        "handler": calculate_average
    },
    {
        "name": "calculate_median",
        "description": "Hitung median dari array angka",
        "parameters": {
            "type": "object",
            "properties": {
                "numbers": {
                    "type": "array",
                    "items": {"type": "number"},
                    "description": "Array angka"
                }
            },
            "required": ["numbers"]
        },
        "handler": calculate_median
    },
    {
        "name": "calculate_standard_deviation",
        "description": "Hitung standar deviasi dari array angka",
        "parameters": {
            "type": "object",
            "properties": {
                "numbers": {
                    "type": "array",
                    "items": {"type": "number"},
                    "description": "Array angka (minimal 2)"
                }
            },
            "required": ["numbers"]
        },
        "handler": calculate_standard_deviation
    },
    {
        "name": "calculate_factorial",
        "description": "Hitung faktorial dari angka (n!)",
        "parameters": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "integer",
                    "description": "Angka yang akan dihitung faktorialnya (0-170)"
                }
            },
            "required": ["number"]
        },
        "handler": calculate_factorial
    },
    {
        "name": "calculate_gcd",
        "description": "Hitung Greatest Common Divisor (GCD/FPB) dari 2 angka",
        "parameters": {
            "type": "object",
            "properties": {
                "a": {
                    "type": "integer",
                    "description": "Angka pertama"
                },
                "b": {
                    "type": "integer",
                    "description": "Angka kedua"
                }
            },
            "required": ["a", "b"]
        },
        "handler": calculate_gcd
    },
    {
        "name": "calculate_lcm",
        "description": "Hitung Least Common Multiple (LCM/KPK) dari 2 angka",
        "parameters": {
            "type": "object",
            "properties": {
                "a": {
                    "type": "integer",
                    "description": "Angka pertama"
                },
                "b": {
                    "type": "integer",
                    "description": "Angka kedua"
                }
            },
            "required": ["a", "b"]
        },
        "handler": calculate_lcm
    },
    {
        "name": "is_prime",
        "description": "Cek apakah angka adalah bilangan prima",
        "parameters": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "integer",
                    "description": "Angka yang akan dicek"
                }
            },
            "required": ["number"]
        },
        "handler": is_prime
    }
]
