package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
	address    = "localhost:50051"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbs.NewImageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	req := pbs.CreateRequest{
		Api: apiVersion,
		Image: &pbs.Image{
			Id:          "123",
			Url:         "blabla",
			Title:       "Hello",
			Description: "Hello Image!",
			Format:      "mime/hello",
		},
	}
	res1, err := c.Create(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	res2, err := c.Read(ctx, &pbs.ReadRequest{Api: apiVersion, Id: "123"})
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)

}
