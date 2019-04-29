package api

import (
	"context"
	"net/http"
	"time"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	st "bitbucket.org/edoardo849/progimage/pkg/storage"

	"github.com/gorilla/mux"
)

const (
	apiVersion = "v1"
)

func handleImageCreate(c pbs.ImageServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		f, fh, err := r.FormFile("image")
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}

		image := st.Image{}
		err = image.Upload(f, fh)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}

		req := pbs.CreateRequest{
			Api: apiVersion,
			Image: &pbs.Image{
				Id:          image.ID,
				Url:         image.URL,
				Title:       image.Title,
				Description: image.Description,
				Format:      image.Format,
			},
		}
		res, err := c.Create(ctx, &req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"id": res.Id})
		return
	}
}

func handleImageGet(c pbs.ImageServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if len(id) == 0 {
			respondWithError(w, http.StatusNotAcceptable, "Missing the id param")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := c.Read(ctx, &pbs.ReadRequest{Api: apiVersion, Id: id})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
		http.Redirect(w, r, res.Image.Url, http.StatusSeeOther)
		return
	}
}
