package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gSt "bitbucket.org/edoardo849/progimage/pkg/protocol/grpc/storage"
)

// Config is configuration for Server
var grpcPort string

func main() {
	gCloudKey := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if gCloudKey != "" {
		log.Printf("Google key file in %s\n", gCloudKey)
		if _, err := os.Stat(gCloudKey); err == nil {
			log.Println("The Google key file exist")
		} else if os.IsNotExist(err) {
			panic("The credentials aren't set")
		}
	}

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
