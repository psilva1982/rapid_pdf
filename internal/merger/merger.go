package merger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// MergePDFs combines multiple PDF files into a single output PDF.
// It uses pdfcpu for reliable, pure-Go PDF merging.
func MergePDFs(inputFiles []string, outputFile string) error {
	if len(inputFiles) == 0 {
		return fmt.Errorf("no input files to merge")
	}

	slog.Info("merging PDFs", "input_count", len(inputFiles), "output", outputFile)

	// Ensure the output directory exists.
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Use pdfcpu's MergeCreateFile to combine all PDFs.
	// The last argument (false) means we don't want to optimize/compress.
	if err := api.MergeCreateFile(inputFiles, outputFile, false, nil); err != nil {
		return fmt.Errorf("failed to merge PDFs: %w", err)
	}

	// Get the size of the merged file for logging.
	info, err := os.Stat(outputFile)
	if err == nil {
		slog.Info("merge complete",
			"output", outputFile,
			"size_mb", fmt.Sprintf("%.2f", float64(info.Size())/(1024*1024)),
		)
	}

	return nil
}

// Cleanup removes the temporary PDF files and their parent directory.
func Cleanup(files []string) {
	if len(files) == 0 {
		return
	}

	// All temp files should be in the same directory.
	tmpDir := filepath.Dir(files[0])

	if err := os.RemoveAll(tmpDir); err != nil {
		slog.Warn("failed to clean up temp directory", "dir", tmpDir, "error", err)
	} else {
		slog.Info("cleaned up temp files", "dir", tmpDir)
	}
}
