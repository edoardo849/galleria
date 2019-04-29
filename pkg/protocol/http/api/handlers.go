package api

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"

	"github.com/gorilla/mux"
)

const (
	apiVersion = "v1"
)

func handleImageCreate(c pbs.StorageServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		f, fh, err := r.FormFile("image")
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		req := pbs.UploadRequest{
			Filename:      fh.Filename,
			Extension:     path.Ext(fh.Filename),
			ContentType:   fh.Header.Get("Content-Type"),
			ContentLength: int64(len(b)),
			Data:          b,
		}
		res, err := c.Upload(context.Background(), &req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}

		respondWithJSON(w, http.StatusOK, res)
		return
	}
}

func handleImageGet(c pbs.StorageServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if len(id) == 0 {
			respondWithError(w, http.StatusNotAcceptable, "Missing the id param")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := c.Read(ctx, &pbs.ReadRequest{Id: id})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}

		w.Header().Set("Content-Type", res.ContentType)
		w.Header().Set("Content-Length", strconv.FormatInt(res.ContentLength, 10))
		if _, err := w.Write(res.Data); err != nil {
			log.Println("unable to write image.")
		}

		return
	}
}
