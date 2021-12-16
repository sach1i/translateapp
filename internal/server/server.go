package server

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
	"translateapp/internal/logger"
	"translateapp/internal/service"
)

type Server struct {
	*mux.Router
	*service.Service
	Logger *zap.Logger
}

func NewServer(logger *zap.Logger) *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		Service: service.NewService(),
		Logger:  logger,
	}
	s.routes()
	return s
}

func (s *Server) Run() http.Server {
	return http.Server{
		Addr:         ":8080",
		Handler:      NewServer(logger.NewLogger()),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
