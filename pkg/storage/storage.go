package storage

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
	Filename    string
	ContentType string
	Extension   string
}
