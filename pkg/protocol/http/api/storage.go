package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"

	"github.com/gorilla/mux"
)

func handleImageCreate(ssc pbs.StorageServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		f, fh, err := r.FormFile("image")
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		req := pbs.UploadRequest{
			Filename:      fh.Filename,
			Extension:     path.Ext(fh.Filename),
			ContentType:   fh.Header.Get("Content-Type"),
			ContentLength: int64(len(b)),
			Data:          b,
		}
		res, err := ssc.Upload(context.Background(), &req)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, res)
		return
	}
}

func handleImageGet(ssc pbs.StorageServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if len(id) == 0 {
			respondWithError(w, http.StatusNotAcceptable, "Missing the id param")
			return
		}

		res, err := imageGetFromCache(ssc, id)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithFile(
			w,
			http.StatusOK,
			res.ContentType,
			strconv.FormatInt(res.ContentLength, 10),
			id,
			res.Data,
		)
		return
	}
}
