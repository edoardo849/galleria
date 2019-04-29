package storage

import (
	"context"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	st "bitbucket.org/edoardo849/progimage/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is the API version provided by the server
	apiVersion = "v1"
)

// ImageService creates a new Image Service
type imageServiceServer struct{}

// checkAPI checks if the API version requested by client is supported by server
func (is *imageServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new Image
func (is imageServiceServer) Create(ctx context.Context, req *pbs.CreateRequest) (*pbs.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := is.checkAPI(req.Api); err != nil {
		return nil, err
	}

	i := &st.Image{
		ID:          req.Image.Id,
		URL:         req.Image.Url,
		Description: req.Image.Description,
		Format:      req.Image.Format,
		Title:       req.Image.Title,
	}
	id, err := st.DB.AddImage(i)

	if err != nil {
		return nil, err
	}

	return &pbs.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

func (is imageServiceServer) Read(ctx context.Context, req *pbs.ReadRequest) (*pbs.ReadResponse, error) {
	// check if the API version requested by client is supported by server
	if err := is.checkAPI(req.Api); err != nil {
		return nil, err
	}

	i, err := st.DB.GetImage(req.Id)

	if err != nil {
		return nil, err
	}

	return &pbs.ReadResponse{
		Api: apiVersion,
		Image: &pbs.Image{
			Id:          i.ID,
			Url:         i.URL,
			Description: i.Description,
			Title:       i.Title,
			Format:      i.Format,
		},
	}, nil

}
