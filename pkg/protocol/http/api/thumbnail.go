package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	pbt "bitbucket.org/edoardo849/progimage/pkg/api/transform"
)

func handleImageThumbnail(dst pbt.ThumbnailServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		f, fh, err := r.FormFile("image")
		width, _ := strconv.ParseInt(r.FormValue("width"), 10, 32)
		height, _ := strconv.ParseInt(r.FormValue("height"), 10, 32)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {

			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := dst.Thumbnail(ctx, &pbt.ThumbnailRequest{
			Data:   b,
			Width:  int32(width),
			Height: int32(height),
		})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		respondWithFile(
			w,
			http.StatusOK,
			res.ContentType,
			strconv.FormatInt(res.ContentLength, 10),
			fmt.Sprintf("%s.%s", fh.Filename, "jpeg"),
			res.Data,
		)
		return
	}
}
