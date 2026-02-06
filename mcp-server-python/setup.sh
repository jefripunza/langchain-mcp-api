#!/bin/bash

# Quick setup script - removes old venv and creates new one with compatible Python

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}🐍 MCP Server Python - Clean Setup${NC}"
echo ""

# Remove old venv if exists
if [ -d "venv" ]; then
    echo -e "${YELLOW}🗑️  Removing old virtual environment...${NC}"
    rm -rf venv
    echo -e "${GREEN}✅ Old venv removed${NC}"
fi

# Find compatible Python version (3.13 or lower)
PYTHON_CMD=""
for cmd in python3.13 python3.12 python3.11 python3.10 python3.9; do
    if command -v $cmd &> /dev/null; then
        VERSION=$($cmd --version 2>&1 | awk '{print $2}')
        MAJOR=$(echo $VERSION | cut -d. -f1)
        MINOR=$(echo $VERSION | cut -d. -f2)
        
        if [ "$MAJOR" -eq 3 ] && [ "$MINOR" -le 13 ] && [ "$MINOR" -ge 9 ]; then
            PYTHON_CMD=$cmd
            echo -e "${GREEN}✅ Found compatible Python: $cmd ($VERSION)${NC}"
            break
        fi
    fi
done

if [ -z "$PYTHON_CMD" ]; then
    echo -e "${RED}❌ No compatible Python found!${NC}"
    echo -e "${YELLOW}⚠️  Pydantic requires Python 3.9 - 3.13${NC}"
    echo -e "${YELLOW}Please install Python 3.13:${NC}"
    echo -e "  ${GREEN}brew install python@3.13${NC}"
    exit 1
fi

# Create new venv
echo ""
echo -e "${YELLOW}📦 Creating virtual environment with $PYTHON_CMD...${NC}"
$PYTHON_CMD -m venv venv

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to create virtual environment${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Virtual environment created${NC}"

# Activate venv
echo ""
echo -e "${YELLOW}🔄 Activating virtual environment...${NC}"
source venv/bin/activate

# Upgrade pip
echo ""
echo -e "${YELLOW}⬆️  Upgrading pip...${NC}"
pip install --upgrade pip

# Install dependencies
echo ""
echo -e "${YELLOW}📦 Installing dependencies...${NC}"
pip install -r requirements.txt

if [ $? -eq 0 ]; then
    echo ""
    echo -e "${GREEN}🎉 Setup completed successfully!${NC}"
    echo ""
    echo -e "${YELLOW}To activate the environment:${NC}"
    echo -e "  ${GREEN}source venv/bin/activate${NC}"
    echo ""
    echo -e "${YELLOW}To run the server:${NC}"
    echo -e "  ${GREEN}./run.sh${NC}"
    echo -e "  ${GREEN}or: python main.py${NC}"
    echo ""
else
    echo ""
    echo -e "${RED}❌ Failed to install dependencies${NC}"
    exit 1
fi
