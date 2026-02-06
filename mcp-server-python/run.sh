#!/bin/bash

# Quick run script for MCP Server Python

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}🐍 Starting MCP Server Python...${NC}"
echo ""

# Check if venv exists
if [ ! -d "venv" ]; then
    echo -e "${RED}❌ Virtual environment not found!${NC}"
    echo -e "${YELLOW}Please run: ./activate.sh${NC}"
    exit 1
fi

# Activate and run
source venv/bin/activate

# Check if dependencies are installed
python -c "import fastapi" 2>/dev/null
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Dependencies not installed!${NC}"
    echo -e "${YELLOW}Please run: ./activate.sh${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Starting server on http://localhost:4050${NC}"
echo ""
python main.py
