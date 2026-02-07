#!/bin/bash

# Script to activate Python virtual environment and install dependencies

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}🐍 MCP Server Python - Environment Setup${NC}"
echo ""

# Find compatible Python version (3.13 or lower)
PYTHON_CMD=""
for cmd in python3.13 python3.12 python3.11 python3.10 python3; do
    if command -v $cmd &> /dev/null; then
        VERSION=$($cmd --version 2>&1 | awk '{print $2}')
        MAJOR=$(echo $VERSION | cut -d. -f1)
        MINOR=$(echo $VERSION | cut -d. -f2)
        
        if [ "$MAJOR" -eq 3 ] && [ "$MINOR" -le 13 ]; then
            PYTHON_CMD=$cmd
            echo -e "${GREEN}✅ Found compatible Python: $cmd ($VERSION)${NC}"
            break
        fi
    fi
done

if [ -z "$PYTHON_CMD" ]; then
    echo -e "${RED}❌ No compatible Python found!${NC}"
    echo -e "${YELLOW}⚠️  Pydantic requires Python 3.13 or lower${NC}"
    echo -e "${YELLOW}Please install Python 3.13:${NC}"
    echo -e "  ${GREEN}brew install python@3.13${NC}"
    exit 1
fi

# Check if venv exists
if [ ! -d "venv" ]; then
    echo -e "${YELLOW}⚠️  Virtual environment not found. Creating with $PYTHON_CMD...${NC}"
    $PYTHON_CMD -m venv venv
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Virtual environment created successfully${NC}"
    else
        echo -e "${RED}❌ Failed to create virtual environment${NC}"
        exit 1
    fi
fi

# Activate virtual environment
echo -e "${YELLOW}🔄 Activating virtual environment...${NC}"
source venv/bin/activate

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Virtual environment activated${NC}"
else
    echo -e "${RED}❌ Failed to activate virtual environment${NC}"
    exit 1
fi

# Check if requirements are installed
echo ""
echo -e "${YELLOW}📦 Checking dependencies...${NC}"
python -c "import fastapi" 2>/dev/null
if [ $? -ne 0 ]; then
    echo -e "${YELLOW}⚠️  Dependencies not installed. Installing...${NC}"
    pip install -r requirements.txt
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Dependencies installed successfully${NC}"
    else
        echo -e "${RED}❌ Failed to install dependencies${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}✅ Dependencies already installed${NC}"
fi

echo ""
echo -e "${GREEN}🎉 Environment ready!${NC}"
echo ""
echo -e "${YELLOW}To run the server:${NC}"
echo -e "  ${GREEN}python main.py${NC}"
echo ""
echo -e "${YELLOW}To deactivate:${NC}"
echo -e "  ${GREEN}deactivate${NC}"
echo ""

# Keep the shell in the virtual environment
exec $SHELL
