package main

import (
	"log"
	"time"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"google.golang.org/grpc"
	"github.com/gorilla/mux"
	"bitbucket.org/edoardo849/progimage/pkg/protocol/http/api"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
	grpcAddr    = "localhost:50051"
	httpAddr = "0.0.0.0:8081"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()


	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM)

	stopServerChan := make(chan struct{})
	server := api.New(
		mux.NewRouter(),
		stopServerChan,
		conn,
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
