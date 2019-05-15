package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"skynet/pkg"
)

type Server struct {
	router *mux.Router
	config *root.ServerConfig
}

func NewServer(u root.UserService, config *root.Config) *Server {
	s := Server{
		router: mux.NewRouter(),
		config: config.Server}

	NewUserRouter(u, s.getSubrouter("/user"))
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port " + s.config.Port)
	if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
