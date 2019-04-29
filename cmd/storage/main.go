package main

import (
	"context"
	"fmt"
	"os"

	gSt "bitbucket.org/edoardo849/progimage/pkg/protocol/grpc/storage"
)

// Config is configuration for Server
var grpcPort string

func main() {
	grpcPort = os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func runServer() error {
	ctx := context.Background()

	if len(grpcPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", grpcPort)
	}

	return gSt.RunServer(ctx, grpcPort)
}
