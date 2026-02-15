package storage

import (
	"context"
	"log/slog"

	"github.com/psilva1982/rapid_pdf/internal/config"
)

// Storage defines the interface for persisting generated PDF files.
// Implementations handle where the file is stored (local disk, S3, etc.)
// and return a URL that can be used to retrieve the file.
type Storage interface {
	// Save persists the given data under the specified filename and returns
	// a URL where the file can be accessed.
	Save(ctx context.Context, filename string, data []byte) (fileURL string, err error)
}

// New creates the appropriate Storage implementation based on configuration.
// If AWS S3 credentials are present in the config, it returns an S3Storage;
// otherwise, it returns a LocalStorage that saves files to ./media.
func New(cfg *config.Config) (Storage, error) {
	if cfg.IsS3Configured() {
		slog.Info("storage backend: AWS S3",
			"bucket", cfg.S3Bucket,
			"region", cfg.S3Region,
		)
		return NewS3Storage(cfg)
	}

	slog.Info("storage backend: local filesystem", "path", "./media")
	return NewLocalStorage("./media")
}
