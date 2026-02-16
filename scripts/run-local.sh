#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
IMAGE_NAME="severos/rapid-pdf"
CONTAINER_NAME="rapid-pdf-local"
PORT="${PORT:-8080}"

echo -e "${GREEN}🐳 Running RapidPDF locally${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Error: Docker is not installed${NC}"
    exit 1
fi

# Check if .env file exists
if [ ! -f ".env" ]; then
    echo -e "${YELLOW}⚠️  Warning: .env file not found${NC}"
    echo "Creating a basic .env file..."
    cat > .env << EOF
PORT=8080
MAX_URLS=10
TIMEOUT_SECONDS=60

# AWS S3 Configuration (optional)
# AWS_S3_BUCKET=
# AWS_S3_REGION=
# AWS_S3_ACCESS_KEY=
# AWS_S3_SECRET_KEY=
EOF
    echo -e "${GREEN}✅ Created .env file${NC}"
    echo
fi

# Load environment variables
export $(grep -v '^#' .env | xargs)

# Stop and remove existing container if running
if [ "$(docker ps -aq -f name=${CONTAINER_NAME})" ]; then
    echo -e "${YELLOW}🛑 Stopping existing container...${NC}"
    docker stop ${CONTAINER_NAME} > /dev/null 2>&1 || true
    docker rm ${CONTAINER_NAME} > /dev/null 2>&1 || true
fi

# Create media directory if it doesn't exist
mkdir -p ./media

# Run the container
echo -e "${GREEN}🚀 Starting container...${NC}"
docker run -d \
    --name ${CONTAINER_NAME} \
    -p ${PORT}:8080 \
    -v "$(pwd)/media:/app/media" \
    -e PORT=${PORT:-8080} \
    -e MAX_URLS=${MAX_URLS:-10} \
    -e TIMEOUT_SECONDS=${TIMEOUT_SECONDS:-60} \
    -e AWS_S3_BUCKET=${AWS_S3_BUCKET:-} \
    -e AWS_S3_REGION=${AWS_S3_REGION:-} \
    -e AWS_S3_ACCESS_KEY=${AWS_S3_ACCESS_KEY:-} \
    -e AWS_S3_SECRET_KEY=${AWS_S3_SECRET_KEY:-} \
    ${IMAGE_NAME}:latest

echo
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo -e "${GREEN}✅ Container started successfully!${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo
echo -e "${BLUE}Container Information:${NC}"
echo "  Name: ${CONTAINER_NAME}"
echo "  Port: ${PORT}"
echo
echo -e "${BLUE}Access URLs:${NC}"
echo "  • API: http://localhost:${PORT}"
echo "  • Swagger: http://localhost:${PORT}/swagger/index.html"
echo
echo -e "${BLUE}View logs:${NC}"
echo "  docker logs -f ${CONTAINER_NAME}"
echo
echo -e "${BLUE}Stop container:${NC}"
echo "  docker stop ${CONTAINER_NAME}"
echo
echo -e "${GREEN}📋 Showing logs (Ctrl+C to exit):${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
docker logs -f ${CONTAINER_NAME}
