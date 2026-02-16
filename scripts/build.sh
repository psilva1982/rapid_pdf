#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
IMAGE_NAME="severos/rapid-pdf"
VERSION="${1:-latest}"

echo -e "${GREEN}🐳 Building RapidPDF Docker Image${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo -e "Image: ${YELLOW}${IMAGE_NAME}:${VERSION}${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Error: Docker is not installed${NC}"
    echo "Please install Docker from https://docs.docker.com/get-docker/"
    exit 1
fi

# Check if Dockerfile exists
if [ ! -f "Dockerfile" ]; then
    echo -e "${RED}❌ Error: Dockerfile not found${NC}"
    echo "Please run this script from the project root directory"
    exit 1
fi

# Build the image
echo -e "${GREEN}🔨 Building Docker image...${NC}"
docker build -t "${IMAGE_NAME}:${VERSION}" .

if [ $? -eq 0 ]; then
    echo
    echo -e "${GREEN}✅ Build successful!${NC}"
    echo
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo -e "${GREEN}Image Information:${NC}"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    docker images "${IMAGE_NAME}:${VERSION}"
    echo
    echo -e "${YELLOW}To run the image:${NC}"
    echo "  docker run -p 8080:8080 ${IMAGE_NAME}:${VERSION}"
    echo
    echo -e "${YELLOW}To run with Docker Compose:${NC}"
    echo "  docker-compose up -d"
    echo
else
    echo
    echo -e "${RED}❌ Build failed${NC}"
    exit 1
fi
