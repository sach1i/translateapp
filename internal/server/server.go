package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"translateapp/internal/service"
)

type Server struct {
	*mux.Router
	*service.Service
}

func NewServer() *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		Service: service.NewService(),
	}
	s.routes()
	return s
}

func (s *Server) Run() http.Server {
	return http.Server{
		Addr:         ":8080",
		Handler:      NewServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
