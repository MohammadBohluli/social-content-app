package main

import (
	"log"

	"github.com/MohammadBohluli/social-content-app/adapter/logger"
	"github.com/MohammadBohluli/social-content-app/config"
	"github.com/MohammadBohluli/social-content-app/delivery/http"

	httpserver "github.com/MohammadBohluli/social-content-app/pkg/http_server"
)

func main() {
	cfg := config.Get()

	logAdapter, err := logger.NewZapLogger()
	if err != nil {
		log.Fatalf("❌Failed to create logger: %v", err)
	}

	server := http.New(cfg, httpserver.New(cfg.HTTPServer), logAdapter)
	server.Serve()

}
