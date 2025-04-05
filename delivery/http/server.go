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
)

type Server struct {
	Cfg        config.Config
	HttpServer pkgHttpServer.Server
	Logger     logger.Logger
}

func New(cfg config.Config, s pkgHttpServer.Server, l logger.Logger) *Server {
	return &Server{
		Cfg:        cfg,
		HttpServer: s,
		Logger:     l,
	}
}

func (s Server) Serve() {
	s.RegisterRoutes()

	go func() {
		if err := s.HttpServer.Start(); err != nil && err != http.ErrServerClosed {
			s.Logger.Fatal("❌ shutting down the server")
		}
	}()

	s.Logger.Info("✅ server up and running", "port", s.Cfg.HTTPServer.Port)

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.Logger.Info("🛑 shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), s.Cfg.HTTPServer.ShutDownCtxTimeout*time.Second)
	defer cancel()

	if err := s.HttpServer.Stop(ctx); err != nil {
		s.Logger.Fatal("❌ server forced to shutdown:", "error", err)
	}

	s.Logger.Info("👋 server exited properly")
}

func (s *Server) RegisterRoutes() {

	// Routes
	v1 := s.HttpServer.Router.Group("/v1")
	v1.GET("/health-check", s.healthCheck)
}
