package api

import (
	"context"
	"net/http"
	"time"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
)

const (
	apiVersion = "v1"
)

func handleImageCreate(c pbs.ImageServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

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
		res, err := c.Create(ctx, &req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
		respondWithJSON(w, http.StatusOK, res)
		return
	}
}

func handleImageGet(c pbs.ImageServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := c.Read(ctx, &pbs.ReadRequest{Api: apiVersion, Id: "123"})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}

		respondWithJSON(w, http.StatusOK, res)
		return
	}
}
