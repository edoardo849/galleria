package decode

import (
	"bytes"
	"context"

	pbd "bitbucket.org/edoardo849/progimage/pkg/api/decode"
	is "bitbucket.org/edoardo849/progimage/pkg/image"
)

// convServiceServer creates a new Image Service
type decodeServiceServer struct{}

func (dss decodeServiceServer) Decode(ctx context.Context, req *pbd.Image) (*pbd.Image, error) {

	r := bytes.NewReader(req.Data)
	img, err := is.Decode(r)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	err = is.Convert(&buf, img, req.Filename)
	if err != nil {
		return nil, err
	}

	return &pbd.Image{
		Data:     buf.Bytes(),
		Filename: req.Filename,
	}, nil
}
