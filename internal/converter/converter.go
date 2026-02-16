package converter

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// ConvertURLToPDF navigates to the given URL using a headless Chrome browser,
// waits for the page to fully load, and saves the rendered page as a PDF.
func ConvertURLToPDF(ctx context.Context, url, outputPath string, timeout time.Duration) error {
	slog.Info("converting URL to PDF", "url", url, "output", outputPath)

	// Create a timeout context for this individual page conversion.
	taskCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var buf []byte
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
		// Wait for the body to be visible (page loaded).
		chromedp.WaitVisible("body", chromedp.ByQuery),
		// Small delay to let async content settle.
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithDisplayHeaderFooter(false).
				// WithHeaderTemplate(fmt.Sprintf(
				// 	`<div style="font-size:8px; width:100%%; text-align:center; color:#888;">%s</div>`, url)).
				// WithFooterTemplate(
				// 	`<div style="font-size:8px; width:100%; text-align:center; color:#888;">
				// 		Página <span class="pageNumber"></span> de <span class="totalPages"></span>
				// 	</div>`).
				WithPaperWidth(8.27).   // A4 width in inches
				WithPaperHeight(11.69). // A4 height in inches
				// WithMarginTop(0.6).
				// WithMarginBottom(0.6).
				// WithMarginLeft(0.4).
				// WithMarginRight(0.4).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to convert %s: %w", url, err)
	}

	if err := os.WriteFile(outputPath, buf, 0644); err != nil {
		return fmt.Errorf("failed to write PDF %s: %w", outputPath, err)
	}

	slog.Info("PDF generated successfully", "url", url, "size_bytes", len(buf))
	return nil
}

// ConvertAll processes a slice of URLs and generates a temporary PDF file for
// each one. It returns the list of generated PDF file paths. The caller is
// responsible for cleaning up the temporary files.
func ConvertAll(ctx context.Context, urls []string, timeout time.Duration) ([]string, error) {
	// Create a temporary directory for intermediate PDFs.
	tmpDir, err := os.MkdirTemp("", "rapid_pdf_*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %w", err)
	}

	slog.Info("starting batch conversion", "url_count", len(urls), "tmp_dir", tmpDir)

	// Create a single browser context to reuse across all pages.
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx,
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("disable-dev-shm-usage", true),
			// Fix for net::ERR_HTTP2_PROTOCOL_ERROR on large pages
			chromedp.Flag("disable-http2", true),
		)...,
	)
	defer allocCancel()

	var pdfPaths []string

	for i, url := range urls {
		// Each URL gets its own browser context (isolated cookies/cache).
		taskCtx, taskCancel := chromedp.NewContext(allocCtx)

		outputPath := filepath.Join(tmpDir, fmt.Sprintf("page_%03d.pdf", i+1))

		if err := ConvertURLToPDF(taskCtx, url, outputPath, timeout); err != nil {
			taskCancel()
			slog.Error("failed to convert URL", "url", url, "error", err)
			return pdfPaths, fmt.Errorf("error on URL #%d (%s): %w", i+1, url, err)
		}

		pdfPaths = append(pdfPaths, outputPath)
		taskCancel()

		slog.Info("progress", "completed", i+1, "total", len(urls))
	}

	return pdfPaths, nil
}
