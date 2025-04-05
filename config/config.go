package config

import httpserver "github.com/MohammadBohluli/social-content-app/pkg/http_server"

type Config struct {
	HTTPServer httpserver.Config `koanf:"http_server"`
}
