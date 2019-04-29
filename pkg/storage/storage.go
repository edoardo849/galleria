package storage

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path"

	"cloud.google.com/go/storage"
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

// Upload uploads the image to the GCP bucket
// and sets the URL in the URL field and the Image ID
func (i *Image) Upload(f multipart.File, fh *multipart.FileHeader) error {
	if StorageBucket == nil {
		return errors.New("storage bucket is missing - check config.go")
	}
	ext := path.Ext(fh.Filename)

	h := sha1.New()
	data := io.TeeReader(f, h)

	name := fmt.Sprintf("%x%s", h.Sum(nil), ext)
	i.ID = name

	// random filename, retaining existing extension.
	// name := uuid.Must(uuid.NewV4()).String() + ext

	ctx := context.Background()
	w := StorageBucket.Object(name).NewWriter(ctx)

	// Warning: storage.AllUsers gives public read access to anyone.
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = fh.Header.Get("Content-Type")
	i.Format = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, data); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"
	i.URL = fmt.Sprintf(publicURL, StorageBucketName, name)

	return nil
}
