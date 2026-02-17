package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	defaultMaxURLs        = 10
	defaultTimeoutSeconds = 60
)

// Config holds the application configuration.
type Config struct {
	MaxURLs             int
	TimeoutSeconds      int
	PageLoadWaitSeconds int
	Port                string

	// S3 storage configuration (optional — if empty, files are saved locally).
	S3Bucket    string
	S3Region    string
	S3AccessKey string
	S3SecretKey string
}

// IsS3Configured returns true when all required S3 environment variables are set.
func (c *Config) IsS3Configured() bool {
	return c.S3Bucket != "" && c.S3Region != "" && c.S3AccessKey != "" && c.S3SecretKey != ""
}

// Load reads the .env file and returns a Config with validated values.
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		// .env is optional; log a warning but don't fail
		fmt.Println("[warn] .env file not found, using defaults")
	}

	maxURLs := defaultMaxURLs
	timeoutSeconds := defaultTimeoutSeconds
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if v := os.Getenv("MAX_URLS"); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("MAX_URLS must be a valid integer: %w", err)
		}
		if parsed < 1 {
			return nil, fmt.Errorf("MAX_URLS must be at least 1, got %d", parsed)
		}
		maxURLs = parsed
	}

	if v := os.Getenv("TIMEOUT_SECONDS"); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("TIMEOUT_SECONDS must be a valid integer: %w", err)
		}
		if parsed < 1 {
			return nil, fmt.Errorf("TIMEOUT_SECONDS must be at least 1, got %d", parsed)
		}
		timeoutSeconds = parsed
	}

	pageLoadWaitSeconds := 5
	if v := os.Getenv("PAGE_LOAD_WAIT_SECONDS"); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("PAGE_LOAD_WAIT_SECONDS must be a valid integer: %w", err)
		}
		if parsed < 0 {
			return nil, fmt.Errorf("PAGE_LOAD_WAIT_SECONDS must be non-negative, got %d", parsed)
		}
		pageLoadWaitSeconds = parsed
	}

	return &Config{
		MaxURLs:             maxURLs,
		TimeoutSeconds:      timeoutSeconds,
		PageLoadWaitSeconds: pageLoadWaitSeconds,
		Port:                port,
		S3Bucket:            os.Getenv("AWS_S3_BUCKET"),
		S3Region:            os.Getenv("AWS_S3_REGION"),
		S3AccessKey:         os.Getenv("AWS_S3_ACCESS_KEY"),
		S3SecretKey:         os.Getenv("AWS_S3_SECRET_KEY"),
	}, nil
}
