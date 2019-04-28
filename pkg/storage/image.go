package storage

import (
	pbStorage "bitbucket.org/edoardo849/progimage/pkg/api/storage"
)

const (
	// apiVersion is the API version provided by the server
	apiVersion = "v1"
)

// ImageDatabase provides thread-safe access to a database of images.
type ImageDatabase interface {
	// ListBooks returns a list of books, ordered by title.
	ListImages() ([]*pbStorage.Image, error)

	// ListImagesCreatedBy returns a list of books, ordered by title, filtered by
	// the user who created the book entry.
	ListImagesCreatedBy(userID string) ([]*pbStorage.Image, error)

	// GetImage retrieves a book by its ID.
	GetImage(id int64) (*pbStorage.Image, error)

	// AddImage saves a given book, assigning it a new ID.
	AddImage(b *pbStorage.Image) (id int64, err error)

	// DeleteImage removes a given book by its ID.
	DeleteImage(id int64) error

	// UpdateImage updates the entry for a given book.
	UpdateImage(b *pbStorage.Image) error

	// Close closes the database, freeing up any available resources.
	// TODO(cbro): Close() should return an error.
	Close()
}

// Image is the image service
type Image struct {
}

// NewImage generates a new image service
func NewImage() Image {
	return Image{}
}

