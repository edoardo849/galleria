package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	pbt "bitbucket.org/edoardo849/progimage/pkg/api/transform"
	"github.com/gorilla/mux"
)

func handleImageConvert(dsc pbt.DecodeServiceClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			filename string
			from     string
			to       string
			res      *pbt.Response
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
