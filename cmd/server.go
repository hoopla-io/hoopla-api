package cmd

import (
	"fmt"
	"github.com/qahvazor/qahvazor/app/config"
	"net/http"
	"time"

	"github.com/qahvazor/qahvazor/app/http/middleware"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(appCfg *config.AppConfig, handler http.Handler) error {
	handler = middleware.CORSMiddleware(handler)

	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", appCfg.HOST, appCfg.PORT),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
