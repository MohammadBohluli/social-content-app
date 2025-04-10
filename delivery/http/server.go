package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MohammadBohluli/social-content-app/adapter/logger"
	"github.com/MohammadBohluli/social-content-app/config"
	pkgHttpServer "github.com/MohammadBohluli/social-content-app/pkg/http_server"
	"github.com/MohammadBohluli/social-content-app/postapp"
	"github.com/MohammadBohluli/social-content-app/repository/psql"
)

type Server struct {
	cfg        config.Config
	httpServer pkgHttpServer.Server
	logger     logger.Logger
	post       postapp.App
}

func New(cfg config.Config, s pkgHttpServer.Server, l logger.Logger, conn *psql.DB) *Server {
	postApp := postapp.New(conn)

	return &Server{
		cfg:        cfg,
		httpServer: s,
		logger:     l,
		post:       postApp,
	}
}

func (s Server) Serve() {
	s.RegisterRoutes()

	go func() {
		if err := s.httpServer.Start(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("❌ shutting down the server")
		}
	}()

	s.logger.Info("✅ server up and running", "port", s.cfg.HTTPServer.Port)

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.logger.Info("🛑 shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.HTTPServer.ShutDownCtxTimeout*time.Second)
	defer cancel()

	if err := s.httpServer.Stop(ctx); err != nil {
		s.logger.Fatal("❌ server forced to shutdown:", "error", err)
	}

	s.logger.Info("👋 server exited properly")
}

func (s *Server) RegisterRoutes() {

	v1 := s.httpServer.Router.Group(config.ApiVersion)
	v1.GET("/health-check", s.healthCheck)

	s.post.Handler.SetRoutes(s.httpServer)

}
