package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Server *http.Server
	Router *mux.Router
}

func NewServer() *Server {
	return &Server{
		Server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
		Router: mux.NewRouter().StrictSlash(true),
	}
}

func (s *Server) Run(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	s.Server.Addr = port
	log.Printf("server starting on %s", port)
	return s.Server.ListenAndServe()
}
