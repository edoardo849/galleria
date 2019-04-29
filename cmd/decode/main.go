package main

import (
	"context"
	"fmt"
	"os"

	gDe "bitbucket.org/edoardo849/progimage/pkg/protocol/grpc/decode"
)

// Config is configuration for Server
var grpcPort string

func main() {
	grpcPort = os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50052"
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

	return gDe.RunServer(ctx, grpcPort)
}
