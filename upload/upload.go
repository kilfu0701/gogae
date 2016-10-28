package upload

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/blobstore"
)

const (
	STORAGE_URL = "https://storage.cloud.google.com/"
	DEFAULT_MAX_SIZE = 1024 * 1024 * 10               // 10MB
)

type Settings struct {
	Bucket  string `validate:"required"` // "your-app.appspot.com"
	Folder  string                       // "my-folder/photo"
	MaxSize int64                        // 1024 * 1024 * 10
	BlobUrl string `validate:"required`  // "/api/photo/upload"
}

// Usage:
//
// import "github.com/kilfu0701/gogae/upload"
//
// settings := upload.Settings{
//      Bucket:  "asd",
//      Folder:  "customer/photo",
//      MaxSize: 1024 * 1024 * 10,    // 10MB
//      BlobUrl: "/api/photo/upload",
// }
//
// url, _ := upload.GenerateUploadURL(ctx, settings)
func GenerateUploadURL(ctx context.Context, settings *Settings) (string, error) {
	if err := validateSettings(settings); err != nil {
		return "", err
	}

	dest := settings.Bucket + "/" + settings.Folder

	option := blobstore.UploadURLOptions{
		MaxUploadBytes: settings.MaxSize,
		StorageBucket: dest,
	}

	url, err := blobstore.UploadURL(ctx, settings.BlobUrl, &option)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func validateSettings(settings *Settings) (error) {
	if settings.Bucket == "" {
		return errors.New("Bucket cannot be empty.")
	}

	if settings.BlobUrl == "" {
		return errors.New("BlobUrl cannot be empty.")
	}

	return nil
}
