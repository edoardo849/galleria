package storage

import (
	pb "bitbucket.org/edoardo849/progimage/pkg/api"
)

const (
	// apiVersion is the API version provided by the server
	apiVersion = "v1"
)

// ImageDatabase provides thread-safe access to a database of images.
type ImageDatabase interface {
	// ListBooks returns a list of books, ordered by title.
	ListImages() ([]*pb.Image, error)

	// ListImagesCreatedBy returns a list of books, ordered by title, filtered by
	// the user who created the book entry.
	ListImagesCreatedBy(userID string) ([]*pb.Image, error)

	// GetImage retrieves a book by its ID.
	GetImage(id int64) (*pb.Image, error)

	// AddImage saves a given book, assigning it a new ID.
	AddImage(b *pb.Image) (id int64, err error)

	// DeleteImage removes a given book by its ID.
	DeleteImage(id int64) error

	// UpdateImage updates the entry for a given book.
	UpdateImage(b *pb.Image) error

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

