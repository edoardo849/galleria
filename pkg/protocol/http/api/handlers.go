package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"time"

	pbd "bitbucket.org/edoardo849/progimage/pkg/api/decode"
	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	st "bitbucket.org/edoardo849/progimage/pkg/storage"

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

func handleImageConvert(dsc pbd.DecodeServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			filename string
			from     string
			to       string
			res      *pbd.DecodeResponse
			err      error
		)

		if r.Method == "GET" {

			vars := mux.Vars(r)
			filename = vars["id"]
			from = vars["from"]
			to = vars["to"]

			if len(filename) == 0 || len(from) == 0 || len(to) == 0 {
				respondWithError(w, http.StatusNotAcceptable, "Missing required params id, from, to")
				return
			}
			res, err = imageConvertFromCache(dsc, fmt.Sprintf("%s.%s", filename, from), from, to)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, err.Error())
				return
			}
			respondWithFile(
				w,
				http.StatusOK,
				res.ContentType,
				strconv.FormatInt(res.ContentLength, 10),
				fmt.Sprintf("%s.%s", filename, to),
				res.Data,
			)
			return

		}

		f, fh, err := r.FormFile("image")
		if r.Method == "POST" && fh != nil {

			fmt.Println("Converting from bytes")
			filename = r.FormValue("filename")
			from = r.FormValue("from")
			to = r.FormValue("to")

			if err != nil {
				respondWithError(w, http.StatusBadRequest, err.Error())
				return
			}

			b, err := ioutil.ReadAll(f)
			if err != nil {

				respondWithError(w, http.StatusBadRequest, err.Error())
				return
			}

			res, err = imageConvertFromBytes(dsc, filename, from, to, b)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, err.Error())
				return
			}

			respondWithFile(
				w,
				http.StatusOK,
				res.ContentType,
				strconv.FormatInt(res.ContentLength, 10),
				fmt.Sprintf("%s.%s", filename, to),
				res.Data,
			)
			return

		}

		var convRequest struct {
			From     string `json:"from"`
			To       string `json:"to"`
			Filename string `json:"filename"`
			URL      string `json:"url"`
		}

		err = json.NewDecoder(r.Body).Decode(&convRequest)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		if len(convRequest.URL) != 0 {

			from = convRequest.From
			to = convRequest.To
			filename = convRequest.Filename

			res, err = imageConvertFromURL(dsc, filename, from, to, convRequest.URL)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, err.Error())
				return
			}
			respondWithFile(
				w,
				http.StatusOK,
				res.ContentType,
				strconv.FormatInt(res.ContentLength, 10),
				fmt.Sprintf("%s.%s", filename, to),
				res.Data,
			)
			return
		}

		respondWithError(w, http.StatusInternalServerError, "Unkwnown error")

		return
	}
}

func imageGetFromCache(ssc pbs.StorageServiceClient, id string) (*pbs.ReadResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return ssc.Get(ctx, &pbs.ReadRequest{Id: id})
}

func imageConvertFromCache(dsc pbd.DecodeServiceClient, filename, from, to string) (*pbd.DecodeResponse, error) {
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", st.StorageBucketName, filename)
	return imageConvertFromURL(dsc, filename, from, to, url)
}

func imageConvertFromBytes(dsc pbd.DecodeServiceClient, filename, from, to string, b []byte) (*pbd.DecodeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return dsc.Decode(ctx, &pbd.DecodeRequest{
		Type:     pbd.DecodeRequest_FROM_BYTES,
		Filename: filename,
		From:     from,
		To:       to,
		Data:     b,
	})
}

func imageConvertFromURL(dsc pbd.DecodeServiceClient, filename, from, to, url string) (*pbd.DecodeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return dsc.Decode(ctx, &pbd.DecodeRequest{
		Type:     pbd.DecodeRequest_FROM_URL,
		Filename: filename,
		From:     from,
		To:       to,
		Url:      url,
	})
}
