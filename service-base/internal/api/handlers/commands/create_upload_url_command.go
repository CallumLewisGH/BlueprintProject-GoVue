package command

import (
	"fmt"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/integrations"
	"github.com/google/uuid"
)

var contentTypeExtensions = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/webp": "webp",
}

// CreateUploadURLCommand returns a signed URL the client can PUT an image to
// directly, plus the URL it will be publicly reachable at afterwards.
//
// Profile pictures use a deterministic per-user path so re-uploading
// overwrites the old one instead of leaking orphaned files.
func CreateUploadURLCommand(userID uuid.UUID, purpose string, contentType string) (uploadURL string, publicURL string, err error) {
	ext, ok := contentTypeExtensions[contentType]
	if !ok {
		return "", "", fmt.Errorf("unsupported content type: %s", contentType)
	}

	var objectKey string
	switch purpose {
	case "profile":
		objectKey = fmt.Sprintf("profile-pictures/%s/avatar.%s", userID, ext)
	default:
		return "", "", fmt.Errorf("unsupported upload purpose: %s", purpose)
	}

	client, err := integrations.GetGCSClient()
	if err != nil {
		return "", "", fmt.Errorf("image storage is not configured: %w", err)
	}

	return client.GenerateUploadURL(objectKey, contentType)
}
