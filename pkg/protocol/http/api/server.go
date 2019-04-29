package api

import (
	"context"
	"log"
	"net/http"
	"os"

	pbs "bitbucket.org/edoardo849/progimage/pkg/api/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// ServerFactory creates a new Server
type ServerFactory func(
	*mux.Router,
	chan struct{},
) *Server

// New Creates a new server
func New(
	r *mux.Router,
	stopChan chan struct{},
	grpcConn *grpc.ClientConn) *Server {

	return &Server{
		router:   mux.NewRouter(),
		stopChan: stopChan,
		grpcConn: grpcConn,
	}
}

// Server is the server
type Server struct {
	router   *mux.Router
	stopChan chan struct{}
	http     *http.Server
	grpcConn *grpc.ClientConn
}

// Run runs the server
func (s *Server) ServeHTTP(http *http.Server) error {
	// Initialize routes
	s.http = http
	// http://zabana.me/notes/enable-cors-in-go-api.html

	s.registerHandlers()
	s.http.Handler = s.router

	// zap.S().Infof("Server listening on %s", s.http.Addr)
	log.Println("Server listening")
	go func() {
		<-s.stopChan
		log.Println("Shutting down server")
		s.http.Shutdown(context.Background())
	}()

	return s.http.ListenAndServe()
}

// Register routes
func (s *Server) registerHandlers() {

	imageClient := pbs.NewImageServiceClient(s.grpcConn)

	// Use gorilla/mux for rich routing.
	// See http://www.gorillatoolkit.org/pkg/mux
	r := s.router.PathPrefix("/v1").Subrouter()

	// Handle all preflight request
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	r.HandleFunc("/status", handleTODO()).Methods("GET")
	r.HandleFunc("/image", handleImageCreate(imageClient)).Methods("POST")
	r.HandleFunc("/image/{id}", handleImageGet(imageClient)).Methods("GET")

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	// [END request_logging]
}
