package storage

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path"

	"cloud.google.com/go/storage"

	uuid "github.com/gofrs/uuid"
)

// ImageDatabase provides thread-safe access to a database of images.
type ImageDatabase interface {
	// ListBooks returns a list of books, ordered by title.
	ListImages() ([]*Image, error)

	// GetImage retrieves a book by its ID.
	GetImage(id string) (*Image, error)

	// AddImage saves a given book, assigning it a new ID.
	AddImage(b *Image) (id string, err error)

	// DeleteImage removes a given book by its ID.
	DeleteImage(id string) error

	// Close closes the database, freeing up any available resources.
	// TODO(cbro): Close() should return an error.
	Close()
}

// Image is the image service
type Image struct {
	ID          string
	URL         string
	Description string
	Title       string
	Format      string
}

// SetID sets the id of the image
func getImageID(u string) string {
	h := sha1.New()
	h.Write([]byte(u))
	return hex.EncodeToString(h.Sum(nil))
}

// UploadToBucket uploads the image to the GCP bucket
// and sets the URL in the URL field and the Image ID
func UploadToBucket(f multipart.File, fh *multipart.FileHeader) (url string, err error) {
	if StorageBucket == nil {
		return "", errors.New("storage bucket is missing - check config.go")
	}
	// random filename, retaining existing extension.
	name := uuid.Must(uuid.NewV4()).String() + path.Ext(fh.Filename)

	ctx := context.Background()
	w := StorageBucket.Object(name).NewWriter(ctx)

	// Warning: storage.AllUsers gives public read access to anyone.
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, f); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"

	return fmt.Sprintf(publicURL, StorageBucketName, name), nil
}
