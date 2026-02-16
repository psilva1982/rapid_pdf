package storage

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// LocalStorage saves PDF files to the local filesystem.
type LocalStorage struct {
	basePath string
}

// NewLocalStorage creates a LocalStorage that persists files under basePath.
// It ensures the directory exists on initialization.
func NewLocalStorage(basePath string) (*LocalStorage, error) {
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create local storage directory %s: %w", basePath, err)
	}

	return &LocalStorage{basePath: basePath}, nil
}

// Save writes the PDF data to a file in the local media directory.
// It generates a unique filename using UUID + timestamp and returns
// a relative URL path like "/media/<filename>.pdf".
func (ls *LocalStorage) Save(_ context.Context, _ string, data []byte) (string, error) {
	filename := generateFilename()
	filePath := filepath.Join(ls.basePath, filename)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", fmt.Errorf("failed to write file to %s: %w", filePath, err)
	}

	fileURL := fmt.Sprintf("/media/%s", filename)

	slog.Info("file saved locally",
		"path", filePath,
		"url", fileURL,
		"size_bytes", len(data),
	)

	return fileURL, nil
}

// generateFilename creates a unique filename with timestamp and UUID.
func generateFilename() string {
	ts := time.Now().Format("20060102_150405")
	id := uuid.New().String()[:8]
	return fmt.Sprintf("%s_%s.pdf", ts, id)
}
