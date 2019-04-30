package storage

import (
	"bytes"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	"bitbucket.org/edoardo849/progimage/pkg/common"
	st "bitbucket.org/edoardo849/progimage/pkg/storage"
	"cloud.google.com/go/storage"
)

// ImageService creates a new Image Service
type storageServiceServer struct{}

// Create new Image
func (ss storageServiceServer) Upload(ctx context.Context, req *pbs.UploadRequest) (*pbs.UploadResponse, error) {

	if st.StorageBucket == nil {
		return nil, errors.New("storage bucket is missing - check config.go")
	}
	r := bytes.NewReader(req.Data)

	h := sha1.New()
	data := io.TeeReader(r, h)

	id := fmt.Sprintf("%x%s", h.Sum(nil), req.Extension)
	log.Printf("Getting image from cache %s\n", id)
	w := st.StorageBucket.Object(id).NewWriter(context.Background())

	// Warning: storage.AllUsers gives public read access to anyone.
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = req.ContentType
	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"

	image := &st.Image{
		ID:            id,
		Filename:      req.Filename,
		URL:           fmt.Sprintf(publicURL, st.StorageBucketName, id),
		ContentType:   req.ContentType,
		ContentLength: req.ContentLength,
		Extension:     req.Extension,
	}

	_, err := st.DB.AddImage(image)
	if err != nil {
		return nil, err
	}

	return &pbs.UploadResponse{
		Id: id,
	}, nil
}

func (ss storageServiceServer) Get(ctx context.Context, req *pbs.ReadRequest) (*pbs.ReadResponse, error) {

	image, err := st.DB.GetImage(req.Id)

	bytes, err := common.GetRawFromURL(image.URL)
	if err != nil {
		return nil, err
	}

	return &pbs.ReadResponse{
		ContentType:   image.ContentType,
		ContentLength: int64(len(bytes)),
		Data:          bytes,
	}, nil

}
