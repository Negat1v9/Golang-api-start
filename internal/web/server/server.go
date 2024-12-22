package server

import (
	"context"
	"log"
	"net/http"

	"github.com/Negat1v9/golang-api-start/internal/config"
)

type Server struct {
	cfg    *config.Config
	server http.Server
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
		server: http.Server{
			Addr: cfg.ServerPort,
		},
	}
}

func (s *Server) Run() error {
	s.initHandlers()

	log.Println("start listen and serve")

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
