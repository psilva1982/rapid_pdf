package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/psilva1982/rapid_pdf/internal/config"
	"github.com/psilva1982/rapid_pdf/internal/converter"
	"github.com/psilva1982/rapid_pdf/internal/merger"
)

// Handler holds the dependencies for the API handlers.
type Handler struct {
	Config *config.Config
}

// NewHandler creates a new Handler with the given configuration.
func NewHandler(cfg *config.Config) *Handler {
	return &Handler{Config: cfg}
}

// GenerateRequest defines the expected JSON body for PDF generation.
type GenerateRequest struct {
	URLs []string `json:"urls" binding:"required,min=1"`
}

// GeneratePDF handles the PDF generation request.
// It accepts a JSON body with a list of URLs, converts them concurrently,
// merges the results, and streams the final PDF back to the client.
//
// @Summary      Generate PDF from URLs
// @Description  Converts a list of URLs to PDF and merges them into a single document.
// @Tags         pdf
// @Accept       json
// @Produce      application/pdf
// @Param        request body GenerateRequest true "List of URLs to convert"
// @Success      200 {file} file "document.pdf"
// @Failure      400 {object} map[string]string "Bad Request"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /generate [post]
func (h *Handler) GeneratePDF(c *gin.Context) {
	var req GenerateRequest

	// bind the JSON body to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate URLs
	for _, u := range req.URLs {
		if u == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty URL provided"})
			return
		}
	}

	if len(req.URLs) > h.Config.MaxURLs {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("too many URLs provided, max is %d", h.Config.MaxURLs)})
		return
	}

	slog.Info("received generate request", "url_count", len(req.URLs))

	// Create a temporary file for the merged PDF
	tmpFile, err := os.CreateTemp("", "rapid_pdf_merged_*.pdf")
	if err != nil {
		slog.Error("failed to create temp file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	// We close the file immediately as we only need the path for now; merger/converter handles IO
	tmpFile.Close()
	outputPath := tmpFile.Name()

	// Ensure cleanup of the final merged file after serving
	defer os.Remove(outputPath)

	// Context for the request is passed down
	ctx := c.Request.Context()

	// 1. Convert all URLs to individual PDFs
	timeout := time.Duration(h.Config.TimeoutSeconds) * time.Second
	pdfFiles, err := converter.ConvertAll(ctx, req.URLs, timeout)
	if err != nil {
		slog.Error("conversion failed", "error", err)
		// Cleanup any partial files
		merger.Cleanup(pdfFiles)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("conversion failed: %v", err)})
		return
	}

	// Ensure intermediate files are cleaned up
	defer merger.Cleanup(pdfFiles)

	// 2. Merge PDFs into the single output file
	if err := merger.MergePDFs(pdfFiles, outputPath); err != nil {
		slog.Error("merge failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("merge failed: %v", err)})
		return
	}

	// 3. Stream the file back to the client
	c.FileAttachment(outputPath, "document.pdf")
}
