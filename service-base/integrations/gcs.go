package integrations

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/storage"
)

type GCSClient struct {
	gcs        *storage.Client
	bucketName string
	saEmail    string
}

var (
	gcsInstance *GCSClient
	gcsInitErr  error
	gcsOnce     sync.Once
)

// GetGCSClient reads Application Default Credentials the same way in every
// environment: a service account key file via GOOGLE_APPLICATION_CREDENTIALS
// locally, or the ambient Cloud Run service identity in prod. Signing needs
// GCS_SERVICE_ACCOUNT_EMAIL set only when running without a private key
// (i.e. in prod) - see the setup notes in this package's README section.
//
// Unlike database.GetDatabase() or auth.NewAuth(), this is invoked lazily on
// the first upload request rather than at startup, so misconfiguration must
// come back as an error to that one request, not log.Fatal - a missing
// bucket name would otherwise take the entire server down on first use.
func GetGCSClient() (*GCSClient, error) {
	gcsOnce.Do(func() {
		ctx := context.Background()

		bucketName := os.Getenv("GCS_BUCKET_NAME")
		if bucketName == "" {
			gcsInitErr = fmt.Errorf("GCS_BUCKET_NAME is not set")
			return
		}

		gcsClient, err := storage.NewClient(ctx)
		if err != nil {
			gcsInitErr = fmt.Errorf("failed to create GCS client: %w", err)
			return
		}

		gcsInstance = &GCSClient{
			gcs:        gcsClient,
			bucketName: bucketName,
			saEmail:    os.Getenv("GCS_SERVICE_ACCOUNT_EMAIL"),
		}
	})
	return gcsInstance, gcsInitErr
}

// GenerateUploadURL returns a short-lived signed URL the caller can PUT the
// object to directly, plus the public URL it will be reachable at once
// uploaded (the bucket must grant allUsers Storage Object Viewer).
func (c *GCSClient) GenerateUploadURL(objectKey string, contentType string) (uploadURL string, publicURL string, err error) {
	opts := &storage.SignedURLOptions{
		Scheme:      storage.SigningSchemeV4,
		Method:      "PUT",
		Expires:     time.Now().Add(15 * time.Minute),
		ContentType: contentType,
	}
	if c.saEmail != "" {
		opts.GoogleAccessID = c.saEmail
	}

	url, err := c.gcs.Bucket(c.bucketName).SignedURL(objectKey, opts)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign upload url: %w", err)
	}

	return url, c.PublicURL(objectKey), nil
}

func (c *GCSClient) PublicURL(objectKey string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", c.bucketName, objectKey)
}
