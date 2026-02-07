from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from typing import Dict, Any, List, Optional
import uvicorn

from tools.encoding import encoding_tools
from tools.hash import hash_tools
from tools.network import network_tools
from tools.math_tools import math_tools
from tools.file_tools import file_tools

app = FastAPI(
    title="MCP Server Python",
    description="Model Context Protocol Server with Python FastAPI",
    version="1.0.0"
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

all_tools = []
all_tools.extend(encoding_tools)
all_tools.extend(hash_tools)
all_tools.extend(network_tools)
all_tools.extend(math_tools)
all_tools.extend(file_tools)

tools_registry = {tool["name"]: tool for tool in all_tools}


class InvokeRequest(BaseModel):
    name: str
    arguments: Dict[str, Any]


@app.get("/")
def root():
    return {
        "message": "🧠 MCP Server Python is running",
        "version": "1.0.0"
    }


@app.get("/health")
def health():
    return {"status": "ok"}


@app.get("/mcp/tools")
def list_tools():
    return [
        {
            "name": tool["name"],
            "description": tool["description"],
            "parameters": tool["parameters"]
        }
        for tool in all_tools
    ]


@app.post("/mcp/invoke")
def invoke_tool(request: InvokeRequest):
    tool = tools_registry.get(request.name)
    
    if not tool:
        raise HTTPException(status_code=404, detail=f"Tool '{request.name}' not found")
    
    try:
        handler = tool["handler"]
        result = handler(request.arguments)
        return {"result": result}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=4050)
