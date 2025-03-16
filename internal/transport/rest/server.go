package rest

import (
	"context"
	"github.com/dnevsky/restaurant-back/internal/pkg/config"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunHttp(handler http.Handler) error {
	var writeTimeout time.Duration
	if config.Config.Debug {
		writeTimeout = time.Second * 120
	} else {
		writeTimeout = config.Config.HTTPConfig.WriteTimeout
	}
	logger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	s.httpServer = &http.Server{
		Addr:              ":" + config.Config.HTTPConfig.Port,
		Handler:           handler,
		ReadTimeout:       config.Config.HTTPConfig.ReadTimeout,
		WriteTimeout:      writeTimeout,
		MaxHeaderBytes:    config.Config.HTTPConfig.MaxHeaderMegabytes << 20,
		ReadHeaderTimeout: config.Config.HTTPConfig.ReadTimeout,
		IdleTimeout:       config.Config.HTTPConfig.ReadTimeout,
		ErrorLog:          logger,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
