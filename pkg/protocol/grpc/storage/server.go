package storage

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	st "bitbucket.org/edoardo849/progimage/pkg/storage"
)

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	pbs.RegisterImageServiceServer(server, imageServiceServer{})

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC image service server...")

			server.GracefulStop()

			// Close the Database
			st.DB.Close()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC image service server...")
	return server.Serve(listen)
}
