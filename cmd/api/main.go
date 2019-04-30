package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	pbt "bitbucket.org/edoardo849/progimage/pkg/api/transform"

	"bitbucket.org/edoardo849/progimage/pkg/protocol/http/api"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {

	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		httpAddr = "0.0.0.0:8081"
	}

	grpcStorageAddr := os.Getenv("GRPC_STORAGE_ADDR")
	if grpcStorageAddr == "" {
		grpcStorageAddr = "localhost:50051"
	}

	grpcDecodeAddr := os.Getenv("GRPC_DECODE_ADDR")
	if grpcDecodeAddr == "" {
		grpcDecodeAddr = "localhost:50052"
	}

	grpcThumbnailAddr := os.Getenv("GRPC_THUMBNAIL_ADDR")
	if grpcThumbnailAddr == "" {
		grpcThumbnailAddr = "localhost:50053"
	}

	// Set up a connection to the storage server.
	stConn, err := grpc.Dial(grpcStorageAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer stConn.Close()
	storageClient := pbs.NewStorageServiceClient(stConn)

	// Set up a connection to the decode server.
	dsConn, err := grpc.Dial(grpcDecodeAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer dsConn.Close()
	decodeClient := pbt.NewDecodeServiceClient(dsConn)

	// Set up a connection to the decode server.
	thConn, err := grpc.Dial(grpcThumbnailAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer dsConn.Close()
	thumbnailClient := pbt.NewThumbnailServiceClient(thConn)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM)

	stopServerChan := make(chan struct{})
	server := api.New(
		mux.NewRouter(),
		stopServerChan,
		storageClient,
		decodeClient,
		thumbnailClient,
	)

	go func() {

		srv := &http.Server{
			Addr: httpAddr,
			// Good practice to set timeouts to avoid Slowloris attacks.
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  5 * time.Second,
		}

		if err := server.ServeHTTP(srv); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}

	}()
	<-stop
	log.Println("shutting down")

	stopServerChan <- struct{}{}
}
