# syntax=docker/dockerfile:1

# ============================================================================
# Build Stage: Compile Go application
# ============================================================================
FROM golang:1.25-alpine AS builder

# Install build dependencies and Chromium (required for chromedp)
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./

# Download dependencies (cached layer)
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimization flags
# -s: Strip debug information
# -w: Strip DWARF symbol table
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -ldflags="-s -w" -o /app/rapid-pdf .

# ============================================================================
# Final Stage: Create minimal runtime image
# ============================================================================
FROM alpine:3.21

# Install runtime dependencies
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    && addgroup -g 1000 appuser \
    && adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/rapid-pdf /app/rapid-pdf

# Create media directory for local PDF storage
RUN mkdir -p /app/media && chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose application port
EXPOSE 8080

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/swagger/index.html || exit 1

# Set environment for Chromium
ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/

# Labels for metadata
LABEL org.opencontainers.image.title="RapidPDF" \
      org.opencontainers.image.description="Efficient Web-to-PDF Converter with AWS S3 support" \
      org.opencontainers.image.version="1.0.0" \
      org.opencontainers.image.authors="Paulo Silva <psilva1982>" \
      org.opencontainers.image.source="https://github.com/psilva1982/rapid_pdf"

# Volume for local PDF storage
VOLUME ["/app/media"]

# Run the application
CMD ["/app/rapid-pdf"]
