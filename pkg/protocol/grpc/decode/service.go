package decode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is the API version provided by the server
	apiVersion = "v1"
)

// convServiceServer creates a new Image Service
type convServiceServer struct{}

// checkAPI checks if the API version requested by client is supported by server
func (cs *convServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}
