#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
DOCKER_USERNAME="severos"
IMAGE_NAME="rapid-pdf"
FULL_IMAGE="${DOCKER_USERNAME}/${IMAGE_NAME}"

echo -e "${GREEN}🚀 Publishing RapidPDF to DockerHub${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Error: Docker is not installed${NC}"
    exit 1
fi

# Check if user is logged in to DockerHub
if ! docker info | grep -q "Username: ${DOCKER_USERNAME}"; then
    echo -e "${YELLOW}⚠️  Not logged in to DockerHub${NC}"
    echo -e "${BLUE}Please login to DockerHub:${NC}"
    docker login
    echo
fi

# Prompt for version
echo -e "${BLUE}Enter the version to publish (e.g., 1.0.0):${NC}"
read -r VERSION

if [ -z "$VERSION" ]; then
    echo -e "${RED}❌ Error: Version cannot be empty${NC}"
    exit 1
fi

# Remove 'v' prefix if present
VERSION=${VERSION#v}

echo
echo -e "${GREEN}📦 Publishing version: ${YELLOW}${VERSION}${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo

# Build the image
echo -e "${GREEN}🔨 Building image...${NC}"
docker build -t "${FULL_IMAGE}:${VERSION}" .

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Build failed${NC}"
    exit 1
fi

# Tag as latest
echo -e "${GREEN}🏷️  Tagging as latest...${NC}"
docker tag "${FULL_IMAGE}:${VERSION}" "${FULL_IMAGE}:latest"

# Push version tag
echo -e "${GREEN}⬆️  Pushing ${FULL_IMAGE}:${VERSION}...${NC}"
docker push "${FULL_IMAGE}:${VERSION}"

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Push failed for version ${VERSION}${NC}"
    exit 1
fi

# Push latest tag
echo -e "${GREEN}⬆️  Pushing ${FULL_IMAGE}:latest...${NC}"
docker push "${FULL_IMAGE}:latest"

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Push failed for latest tag${NC}"
    exit 1
fi

echo
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo -e "${GREEN}✅ Successfully published to DockerHub!${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo
echo -e "${BLUE}Published tags:${NC}"
echo "  • ${FULL_IMAGE}:${VERSION}"
echo "  • ${FULL_IMAGE}:latest"
echo
echo -e "${BLUE}Users can pull your image with:${NC}"
echo "  docker pull ${FULL_IMAGE}:${VERSION}"
echo "  docker pull ${FULL_IMAGE}:latest"
echo
echo -e "${BLUE}View on DockerHub:${NC}"
echo "  https://hub.docker.com/r/${DOCKER_USERNAME}/${IMAGE_NAME}"
echo
