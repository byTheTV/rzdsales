package server

import (
	"net/http"
	"rzd-sales/backend/internal/config"
)

type Server struct {
	*http.Server
}

func New(cfg config.ServerConfig, handler http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}
