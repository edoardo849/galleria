package transform

import (
	"bytes"
	"context"
	"fmt"
	"log"

	tf "bitbucket.org/edoardo849/progimage/pkg/api/transform"
	"bitbucket.org/edoardo849/progimage/pkg/common"
	im "bitbucket.org/edoardo849/progimage/pkg/image"
)

// convServiceServer creates a new Image Service
type decodeServiceServer struct{}

func (ds decodeServiceServer) Decode(ctx context.Context, req *tf.DecodeRequest) (*tf.Response, error) {
	log.Printf("Decoding image from %s to %s of type %s", req.From, req.To, req.Type)

	var b []byte
	var err error
	switch dType := req.Type; dType {
	case tf.DecodeRequest_FROM_URL:
		b, err = common.GetRawFromURL(req.Url)
		if err != nil {
			return nil, err
		}
		log.Printf("Downloaded image from %s", req.Url)
	case tf.DecodeRequest_FROM_BYTES:
		b = req.Data
		log.Println("Received bytes")
	}

	r := bytes.NewReader(b)
	img, err := im.Decode(r)
	if err != nil {
		log.Printf("Error while decoding: %v", err)
		return nil, err
	}
	var buf bytes.Buffer

	err = im.Convert(&buf, img, fmt.Sprintf("%s.%s", req.Filename, req.To))
	if err != nil {
		log.Printf("Error while converting: %v", err)
		return nil, err
	}
	bb := buf.Bytes()
	return &tf.Response{
		Data:          bb,
		ContentType:   fmt.Sprintf("image/%s", req.To),
		ContentLength: int64(len(bb)),
	}, nil
}
