package storage

import (
	"errors"
	"fmt"
	"sync"
)

// Ensure memoryDB conforms to the ImageDatabase interface.
var _ ImageDatabase = &memoryDB{}

// memoryDB is a simple in-memory persistence layer for images.
type memoryDB struct {
	mu     sync.Mutex
	images map[string]*Image // maps from image ID to Image.
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		images: make(map[string]*Image),
	}
}

// ListImages returns a list of images
func (db *memoryDB) ListImages() ([]*Image, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var images []*Image
	for _, i := range db.images {
		images = append(images, i)
	}

	return images, nil
}

// Close closes the database.
func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.images = nil
}

// GetIMage retrieves an image by its ID.
func (db *memoryDB) GetImage(id string) (*Image, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	image, ok := db.images[id]
	if !ok {
		return nil, fmt.Errorf("memorydb: image not found with ID %s", id)
	}
	return image, nil
}

// AddImage saves a given image, assigning it a new ID.
func (db *memoryDB) AddImage(i *Image) (id string, err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.images[i.ID] = i
	return i.ID, nil
}

// DeleteBook removes a given book by its ID.
func (db *memoryDB) DeleteImage(id string) error {
	if id == "" {
		return errors.New("memorydb: image with unassigned ID passed into deleteImage")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.images[id]; !ok {
		return fmt.Errorf("memorydb: could not delete image with ID %s, does not exist", id)
	}
	delete(db.images, id)
	return nil
}
