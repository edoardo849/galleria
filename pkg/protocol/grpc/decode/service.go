package decode

import (
	"bytes"
	"context"
	"fmt"
	"log"

	pbd "bitbucket.org/edoardo849/progimage/pkg/api/decode"
	"bitbucket.org/edoardo849/progimage/pkg/common"
	im "bitbucket.org/edoardo849/progimage/pkg/image"
)

// convServiceServer creates a new Image Service
type decodeServiceServer struct{}

func (dss decodeServiceServer) Decode(ctx context.Context, req *pbd.DecodeRequest) (*pbd.DecodeResponse, error) {
	log.Printf("Decoding image from %s to %s of type %s", req.From, req.To, req.Type)

	var b []byte
	var err error
	switch dType := req.Type; dType {
	case pbd.DecodeRequest_FROM_URL:
		b, err = common.GetRawFromURL(req.Url)
		if err != nil {
			return nil, err
		}
		log.Printf("Downloaded image from %s", req.Url)
	case pbd.DecodeRequest_FROM_BYTES:
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
	return &pbd.DecodeResponse{
		Data:          bb,
		ContentType:   fmt.Sprintf("image/%s", req.To),
		ContentLength: int64(len(bb)),
	}, nil
}
