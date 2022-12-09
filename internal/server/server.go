package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (srv *Server) Run(port string) error {
	srv.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           nil,
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	return srv.httpServer.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
