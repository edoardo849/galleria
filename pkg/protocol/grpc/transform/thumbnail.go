package transform

import (
	"bytes"
	"context"
	"log"

	tf "bitbucket.org/edoardo849/progimage/pkg/api/transform"
	im "bitbucket.org/edoardo849/progimage/pkg/image"
)

// convServiceServer creates a new Image Service
type thumbnailServiceServer struct{}

func (ts thumbnailServiceServer) Thumbnail(ctx context.Context, req *tf.ThumbnailRequest) (*tf.Response, error) {
	log.Printf("Thumbnail from %dx%d", req.Width, req.Height)

	r := bytes.NewReader(req.Data)
	img, err := im.Decode(r)
	if err != nil {
		log.Printf("Error while decoding: %v", err)
		return nil, err
	}
	var buf bytes.Buffer

	err = im.Thumbnail(&buf, img, int(req.Width), int(req.Height))
	if err != nil {
		log.Printf("Error while converting: %v", err)
		return nil, err
	}
	bb := buf.Bytes()
	return &tf.Response{
		Data:          bb,
		ContentType:   "image/jpeg",
		ContentLength: int64(len(bb)),
	}, nil
}
