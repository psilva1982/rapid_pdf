package storage

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/psilva1982/rapid_pdf/internal/config"
)

// S3Storage uploads PDF files to an AWS S3 bucket.
type S3Storage struct {
	client *s3.Client
	bucket string
	region string
}

// NewS3Storage creates an S3Storage using the credentials from config.
func NewS3Storage(cfg *config.Config) (*S3Storage, error) {
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(cfg.S3Region),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.S3AccessKey,
				cfg.S3SecretKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg)

	return &S3Storage{
		client: client,
		bucket: cfg.S3Bucket,
		region: cfg.S3Region,
	}, nil
}

// Save uploads the PDF data to S3 and returns the public URL of the object.
func (ss *S3Storage) Save(ctx context.Context, _ string, data []byte) (string, error) {
	key := generateS3Key()

	input := &s3.PutObjectInput{
		Bucket:      aws.String(ss.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/pdf"),
	}

	if _, err := ss.client.PutObject(ctx, input); err != nil {
		return "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", ss.bucket, ss.region, key)

	slog.Info("file uploaded to S3",
		"bucket", ss.bucket,
		"key", key,
		"url", fileURL,
		"size_bytes", len(data),
	)

	return fileURL, nil
}

// generateS3Key creates a unique S3 object key with date prefix for organization.
func generateS3Key() string {
	now := time.Now()
	datePrefix := now.Format("2006/01/02")
	ts := now.Format("150405")
	id := uuid.New().String()[:8]
	return fmt.Sprintf("pdfs/%s/%s_%s.pdf", datePrefix, ts, id)
}
