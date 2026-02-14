package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/psilva1982/rapid_pdf/docs"
	"github.com/psilva1982/rapid_pdf/internal/api"
	"github.com/psilva1982/rapid_pdf/internal/config"
	"github.com/psilva1982/rapid_pdf/internal/converter"
	"github.com/psilva1982/rapid_pdf/internal/merger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const defaultOutputFile = "output.pdf"

// @title           RapidPDF API
// @version         1.0
// @description     Efficient Web-to-PDF Converter API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	// Set up structured logging.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	slog.Info("🚀 RapidPDF — Web-to-PDF Converter")

	// Load configuration from .env.
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Check if running in CLI mode (arguments provided) or Server mode (no arguments)
	args := os.Args[1:]
	if len(args) == 0 {
		runServer(cfg)
	} else {
		runCLI(cfg, args)
	}
}

func runServer(cfg *config.Config) {
	slog.Info("Starting server mode on :8080")

	// Initialize API handler with configuration
	handler := api.NewHandler(cfg)

	r := gin.Default()

	r.POST("/generate", handler.GeneratePDF)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

func runCLI(cfg *config.Config, urls []string) {
	// Validate number of URLs.
	if len(urls) > cfg.MaxURLs {
		slog.Error("too many URLs provided",
			"provided", len(urls),
			"max_allowed", cfg.MaxURLs,
		)
		fmt.Printf("\n❌ Error: you provided %d URLs, but the maximum allowed is %d.\n", len(urls), cfg.MaxURLs)
		fmt.Println("   Adjust MAX_URLS in .env to increase the limit.")
		os.Exit(1)
	}

	// Validate each URL format.
	for i, u := range urls {
		if !isValidURL(u) {
			slog.Error("invalid URL", "index", i+1, "url", u)
			fmt.Printf("\n❌ Error: invalid URL #%d: %s\n", i+1, u)
			fmt.Println("   URLs must start with http:// or https://")
			os.Exit(1)
		}
	}

	slog.Info("configuration loaded",
		"max_urls", cfg.MaxURLs,
		"urls_provided", len(urls),
	)

	// Start conversion.
	start := time.Now()
	ctx := context.Background()

	fmt.Println()
	fmt.Printf("📄 Converting %d %s to PDF...\n", len(urls), pluralize(len(urls), "page", "pages"))
	fmt.Println(strings.Repeat("─", 50))

	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second
	pdfFiles, err := converter.ConvertAll(ctx, urls, timeout)
	if err != nil {
		slog.Error("conversion failed", "error", err)
		merger.Cleanup(pdfFiles) // Clean up any partial results.
		fmt.Printf("\n❌ Conversion failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(strings.Repeat("─", 50))
	fmt.Printf("✅ All pages converted. Merging into %s...\n", defaultOutputFile)

	// Merge all PDFs into one.
	if err := merger.MergePDFs(pdfFiles, defaultOutputFile); err != nil {
		slog.Error("merge failed", "error", err)
		merger.Cleanup(pdfFiles)
		fmt.Printf("\n❌ Merge failed: %v\n", err)
		os.Exit(1)
	}

	// Clean up temporary files.
	merger.Cleanup(pdfFiles)

	elapsed := time.Since(start)
	fmt.Println()
	fmt.Printf("🎉 Done! PDF saved as: %s\n", defaultOutputFile)
	fmt.Printf("⏱  Completed in %s\n", elapsed.Round(time.Millisecond))
	fmt.Println()
}

// isValidURL checks if the given string is a valid HTTP/HTTPS URL.
func isValidURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

// pluralize returns singular or plural form based on count.
func pluralize(count int, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}
