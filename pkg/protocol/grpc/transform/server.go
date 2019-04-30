package transform

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	tf "bitbucket.org/edoardo849/progimage/pkg/api/transform"
)

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	tf.RegisterDecodeServiceServer(server, decodeServiceServer{})
	tf.RegisterThumbnailServiceServer(server, thumbnailServiceServer{})

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC transform service server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC transform service server...")
	return server.Serve(listen)
}
