package config

import (
	httpserver "github.com/MohammadBohluli/social-content-app/pkg/http_server"
	"github.com/MohammadBohluli/social-content-app/repository/psql"
)

type Config struct {
	HTTPServer httpserver.Config `koanf:"http_server"`
	PSQL       psql.Config       `koanf:"psql"`
}
