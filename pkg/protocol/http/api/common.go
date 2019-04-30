package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	pbt "bitbucket.org/edoardo849/progimage/pkg/api/transform"
	st "bitbucket.org/edoardo849/progimage/pkg/storage"
)

func imageGetFromCache(ssc pbs.StorageServiceClient, id string) (*pbs.ReadResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return ssc.Get(ctx, &pbs.ReadRequest{Id: id})
}

func imageConvertFromCache(dsc pbt.DecodeServiceClient, filename, from, to string) (*pbt.Response, error) {
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", st.StorageBucketName, filename)
	return imageConvertFromURL(dsc, filename, from, to, url)
}

func imageConvertFromBytes(dsc pbt.DecodeServiceClient, filename, from, to string, b []byte) (*pbt.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return dsc.Decode(ctx, &pbt.DecodeRequest{
		Type:     pbt.DecodeRequest_FROM_BYTES,
		Filename: filename,
		From:     from,
		To:       to,
		Data:     b,
	})
}

func imageConvertFromURL(dsc pbt.DecodeServiceClient, filename, from, to, url string) (*pbt.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return dsc.Decode(ctx, &pbt.DecodeRequest{
		Type:     pbt.DecodeRequest_FROM_URL,
		Filename: filename,
		From:     from,
		To:       to,
		Url:      url,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSONString(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}

func respondWithFile(w http.ResponseWriter, code int, contentType string, contentLength string, fileName string, payload []byte) {

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", contentLength)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
	w.WriteHeader(code)

	if _, err := w.Write(payload); err != nil {
		log.Println("unable to write payload.")
	}
}

func withBasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			respondWithError(w, http.StatusUnauthorized, "Not authorized")
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Not authorized")
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			respondWithError(w, http.StatusUnauthorized, "Not authorized")
			return
		}
		//dXNlcm5hbWU6cGFzc3dvcmQ=
		if pair[0] != "username" || pair[1] != "password" {
			respondWithError(w, http.StatusUnauthorized, "Not authorized")
			return
		}

		h(w, r)
	}
}

func handleTODO() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, map[string]string{"data": "TODO"})
		return
	}
}

func handle404() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, http.StatusNotFound, "Not found")
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h(w, r)
	}

}
