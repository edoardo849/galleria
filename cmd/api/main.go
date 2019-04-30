package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pbd "bitbucket.org/edoardo849/progimage/pkg/api/decode"
	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"

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

	// Set up a connection to the storage server.
	stConn, err := grpc.Dial(grpcStorageAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer stConn.Close()
	storageClient := pbs.NewStorageServiceClient(stConn)

	// Set up a connection to the storage server.
	dsConn, err := grpc.Dial(grpcDecodeAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer dsConn.Close()
	decodeClient := pbd.NewDecodeServiceClient(dsConn)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM)

	stopServerChan := make(chan struct{})
	server := api.New(
		mux.NewRouter(),
		stopServerChan,
		storageClient,
		decodeClient,
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
