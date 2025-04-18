package main

import (
	"log"

	"github.com/MohammadBohluli/social-content-app/adapter/logger"
	"github.com/MohammadBohluli/social-content-app/config"
	"github.com/MohammadBohluli/social-content-app/delivery/http"
	"github.com/MohammadBohluli/social-content-app/repository/psql"

	httpserver "github.com/MohammadBohluli/social-content-app/pkg/http_server"
	"github.com/MohammadBohluli/social-content-app/pkg/validator"
)

func main() {
	cfg := config.Get()

	// create log adapter
	logAdapter, err := logger.NewZapLogger()
	if err != nil {
		log.Fatalf("❌ failed to create logger: %v", err)
	}

	// create validator
	validate := validator.New()

	// connect to postgreSQL
	db := psql.New(cfg.PSQL, logAdapter)
	db.Conn()

	// create server
	server := http.New(cfg, httpserver.New(cfg.HTTPServer), logAdapter, db, validate)
	server.Serve()

}
