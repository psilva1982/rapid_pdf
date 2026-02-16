# 🐳 Docker Guide for RapidPDF

Complete guide for using RapidPDF with Docker, including local development, production deployment, and publishing to DockerHub.

## Table of Contents

- [Quick Start](#quick-start)
- [Docker Images](#docker-images)
- [Configuration](#configuration)
- [Development](#development)
- [Production Deployment](#production-deployment)
- [Building and Publishing](#building-and-publishing)
- [Troubleshooting](#troubleshooting)

---

## Quick Start

### Pull and Run from DockerHub

```bash
# Pull the latest version
docker pull severos/rapid-pdf:latest

# Run the container
docker run -d \
  --name rapid-pdf \
  -p 8080:8080 \
  -v $(pwd)/media:/app/media \
  severos/rapid-pdf:latest
```

Access the API at `http://localhost:8080` and Swagger UI at `http://localhost:8080/swagger/index.html`.

### Using Docker Compose

```bash
# Create .env file with your configuration
cp .env.example .env

# Start the service
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the service
docker-compose down
```

---

## Docker Images

### Official Images

Available on DockerHub: [`severos/rapid-pdf`](https://hub.docker.com/r/severos/rapid-pdf)

**Tags:**

- `latest` - Latest stable release
- `v1.0.0` - Specific version

**Image Details:**

- **Base**: Alpine Linux 3.21
- **Size**: ~100-150MB (optimized multi-stage build)
- **Architecture**: linux/amd64
- **Go Version**: 1.25
- **Includes**: Chromium for PDF rendering

---

## Configuration

### Environment Variables

| Variable            | Description                          | Default | Required |
| :------------------ | :----------------------------------- | :------ | :------- |
| `PORT`              | Server port                          | `8080`  | No       |
| `MAX_URLS`          | Maximum URLs per request             | `10`    | No       |
| `TIMEOUT_SECONDS`   | Rendering timeout per page (seconds) | `60`    | No       |
| `AWS_S3_BUCKET`     | S3 bucket name for PDF storage       | -       | No       |
| `AWS_S3_REGION`     | AWS region (e.g., `us-east-1`)       | -       | No       |
| `AWS_S3_ACCESS_KEY` | AWS access key                       | -       | No       |
| `AWS_S3_SECRET_KEY` | AWS secret key                       | -       | No       |

### Example .env File

```env
PORT=8080
MAX_URLS=10
TIMEOUT_SECONDS=60

# Optional: AWS S3 Configuration
AWS_S3_BUCKET=my-pdf-bucket
AWS_S3_REGION=us-east-1
AWS_S3_ACCESS_KEY=AKIAIOSFODNN7EXAMPLE
AWS_S3_SECRET_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

### Volume Mounts

**Local PDF Storage:**

```bash
-v $(pwd)/media:/app/media
```

This mounts your local `media` directory to store generated PDFs when S3 is not configured.

---

## Development

### Building Locally

```bash
# Using the build script
./scripts/build.sh v1.0.0

# Or manually
docker build -t severos/rapid-pdf:v1.0.0 .
```

### Running Locally

```bash
# Using the run script (recommended)
./scripts/run-local.sh

# Or manually
docker run -d \
  --name rapid-pdf-local \
  -p 8080:8080 \
  -v $(pwd)/media:/app/media \
  --env-file .env \
  severos/rapid-pdf:latest
```

### Development with Docker Compose

```yaml
version: "3.8"

services:
  rapid-pdf:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    volumes:
      - ./media:/app/media
    env_file:
      - .env
    restart: unless-stopped
```

Start development environment:

```bash
docker-compose up --build
```

---

## Production Deployment

### Docker Run

```bash
docker run -d \
  --name rapid-pdf \
  --restart unless-stopped \
  -p 8080:8080 \
  -v /var/lib/rapid-pdf/media:/app/media \
  -e PORT=8080 \
  -e MAX_URLS=20 \
  -e TIMEOUT_SECONDS=90 \
  -e AWS_S3_BUCKET=prod-pdf-bucket \
  -e AWS_S3_REGION=us-east-1 \
  -e AWS_S3_ACCESS_KEY=${AWS_ACCESS_KEY} \
  -e AWS_S3_SECRET_KEY=${AWS_SECRET_KEY} \
  --memory="2g" \
  --cpus="1.5" \
  severos/rapid-pdf:latest
```

### Docker Compose (Production)

```yaml
version: "3.8"

services:
  rapid-pdf:
    image: severos/rapid-pdf:latest
    container_name: rapid-pdf-prod
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - MAX_URLS=20
      - TIMEOUT_SECONDS=90
      - AWS_S3_BUCKET=${AWS_S3_BUCKET}
      - AWS_S3_REGION=${AWS_S3_REGION}
      - AWS_S3_ACCESS_KEY=${AWS_S3_ACCESS_KEY}
      - AWS_S3_SECRET_KEY=${AWS_S3_SECRET_KEY}
    volumes:
      - rapid-pdf-data:/app/media
    restart: always
    deploy:
      resources:
        limits:
          memory: 2G
          cpus: "1.5"
    healthcheck:
      test:
        [
          "CMD",
          "wget",
          "--spider",
          "-q",
          "http://localhost:8080/swagger/index.html",
        ]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
    networks:
      - web

volumes:
  rapid-pdf-data:

networks:
  web:
    external: true
```

### Behind Reverse Proxy (Nginx)

```nginx
server {
    listen 80;
    server_name pdf.example.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # Increase timeout for PDF generation
        proxy_read_timeout 120s;
        proxy_connect_timeout 120s;
    }
}
```

---

## Building and Publishing

### Prerequisites

```bash
# Login to DockerHub
docker login
```

### Build Local Image

```bash
# Using the automated script
./scripts/build.sh v1.0.0

# Or manually
docker build -t severos/rapid-pdf:v1.0.0 .
docker tag severos/rapid-pdf:v1.0.0 severos/rapid-pdf:latest
```

### Publish to DockerHub

```bash
# Using the automated script (recommended)
./scripts/publish.sh

# Or manually
docker push severos/rapid-pdf:v1.0.0
docker push severos/rapid-pdf:latest
```

### Multi-Platform Build (Advanced)

For building images for multiple architectures:

```bash
# Create and use buildx builder
docker buildx create --name multiplatform --use

# Build for multiple platforms
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t severos/rapid-pdf:v1.0.0 \
  -t severos/rapid-pdf:latest \
  --push \
  .
```

---

## Troubleshooting

### Container Won't Start

```bash
# Check container logs
docker logs rapid-pdf

# Check container status
docker ps -a | grep rapid-pdf

# Inspect container
docker inspect rapid-pdf
```

### Chromium Issues

If you encounter Chromium-related errors:

```bash
# Run container with additional debugging
docker run -it --rm \
  -e CHROME_BIN=/usr/bin/chromium-browser \
  severos/rapid-pdf:latest \
  /bin/sh

# Test Chromium inside container
chromium-browser --version
```

### Memory Issues

PDF generation can be memory-intensive. Increase container memory:

```bash
docker run -d \
  --name rapid-pdf \
  --memory="4g" \
  --memory-swap="4g" \
  severos/rapid-pdf:latest
```

Or in `docker-compose.yml`:

```yaml
deploy:
  resources:
    limits:
      memory: 4G
```

### Permission Issues

If you have permission issues with the media volume:

```bash
# Fix permissions on host
chmod -R 777 ./media

# Or use a named volume
docker volume create rapid-pdf-media
docker run -v rapid-pdf-media:/app/media severos/rapid-pdf:latest
```

### Health Check Failures

If health checks are failing:

```bash
# Test health endpoint manually
docker exec rapid-pdf wget --spider -q http://localhost:8080/swagger/index.html

# Disable health check temporarily
docker run --no-healthcheck severos/rapid-pdf:latest
```

### S3 Connection Issues

For AWS S3 connectivity problems:

```bash
# Verify AWS credentials
docker exec rapid-pdf env | grep AWS

# Test S3 access from container
docker exec -it rapid-pdf /bin/sh
# Inside container:
wget https://s3.amazonaws.com
```

---

## Performance Tuning

### Optimize for High Load

```yaml
services:
  rapid-pdf:
    image: severos/rapid-pdf:latest
    deploy:
      replicas: 3
      resources:
        limits:
          cpus: "2"
          memory: 3G
        reservations:
          cpus: "1"
          memory: 1G
    environment:
      - MAX_URLS=5
      - TIMEOUT_SECONDS=45
```

### Caching Layers

To improve build times, leverage Docker layer caching:

```dockerfile
# Dependencies are cached separately from code
COPY go.mod go.sum ./
RUN go mod download

# Code changes don't invalidate dependency cache
COPY . .
RUN go build ...
```

---

## Security Best Practices

1. **Use Non-Root User**: The image runs as `appuser` (UID 1000)
2. **Read-Only Root Filesystem** (optional):
   ```bash
   docker run --read-only -v /app/media:rw severos/rapid-pdf:latest
   ```
3. **Secrets Management**: Use Docker secrets or environment variables
4. **Network Isolation**: Run on isolated Docker networks
5. **Regular Updates**: Keep the base image and dependencies updated

---

## Integration Examples

### Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rapid-pdf
spec:
  replicas: 3
  selector:
    matchLabels:
      app: rapid-pdf
  template:
    metadata:
      labels:
        app: rapid-pdf
    spec:
      containers:
        - name: rapid-pdf
          image: severos/rapid-pdf:latest
          ports:
            - containerPort: 8080
          env:
            - name: MAX_URLS
              value: "10"
            - name: AWS_S3_BUCKET
              valueFrom:
                secretKeyRef:
                  name: rapid-pdf-secrets
                  key: s3-bucket
          resources:
            limits:
              memory: "2Gi"
              cpu: "1000m"
            requests:
              memory: "1Gi"
              cpu: "500m"
```

### Docker Swarm

```yaml
version: "3.8"

services:
  rapid-pdf:
    image: severos/rapid-pdf:latest
    deploy:
      replicas: 5
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure
    ports:
      - "8080:8080"
    networks:
      - overlay-net

networks:
  overlay-net:
    driver: overlay
```

---

## License

MIT © [Paulo Silva](https://github.com/psilva1982)

For more information, see the [main README](README.md).
