package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	pbt "bitbucket.org/edoardo849/progimage/pkg/api/transform"

	"github.com/gorilla/mux"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// New Creates a new server
func New(
	r *mux.Router,
	stopChan chan struct{},
	ssc pbs.StorageServiceClient,
	dsc pbt.DecodeServiceClient,
	tsc pbt.ThumbnailServiceClient,

) *Server {

	return &Server{
		router:   mux.NewRouter(),
		stopChan: stopChan,
		ssc:      ssc,
		dsc:      dsc,
		tsc:      tsc,
	}
}

// Server is the server
type Server struct {
	router   *mux.Router
	stopChan chan struct{}
	http     *http.Server
	ssc      pbs.StorageServiceClient
	dsc      pbt.DecodeServiceClient
	tsc      pbt.ThumbnailServiceClient
}

// Run runs the server
func (s *Server) ServeHTTP(http *http.Server) error {
	// Initialize routes
	s.http = http
	// http://zabana.me/notes/enable-cors-in-go-api.html

	s.registerHandlers()
	s.http.Handler = s.router
	log.Printf("Server listening on %s\n", s.http.Addr)

	go func() {
		<-s.stopChan
		log.Println("Shutting down server")
		s.http.Shutdown(context.Background())
	}()

	return s.http.ListenAndServe()
}

// Register routes
func (s *Server) registerHandlers() {

	// Use gorilla/mux for rich routing.
	// See http://www.gorillatoolkit.org/pkg/mux
	r := s.router.PathPrefix(fmt.Sprintf("/%s", apiVersion)).Subrouter()

	// Handle all preflight request
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	r.HandleFunc("/image", handleImageCreate(s.ssc)).Methods("POST")

	// The Router will first try to match the image converter: do not change
	// this order
	r.HandleFunc("/image/{id}.{from}.{to}", handleImageConvert(s.dsc)).Methods("GET")
	r.HandleFunc("/image/{id}", handleImageGet(s.ssc)).Methods("GET")

	r.HandleFunc("/image/convert", handleImageConvert(s.dsc)).Methods("POST")
	r.HandleFunc("/image/thumbnail", handleImageThumbnail(s.tsc)).Methods("POST")

}
