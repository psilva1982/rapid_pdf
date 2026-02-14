package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const defaultMaxURLs = 10

// Config holds the application configuration.
type Config struct {
	MaxURLs int
}

// Load reads the .env file and returns a Config with validated values.
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		// .env is optional; log a warning but don't fail
		fmt.Println("[warn] .env file not found, using defaults")
	}

	maxURLs := defaultMaxURLs

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

	return &Config{MaxURLs: maxURLs}, nil
}
