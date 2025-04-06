package httpserver

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router *echo.Echo
	Config Config
}

type Config struct {
	Port               int           `koanf:"port"`
	CORS               Cors          `koanf:"cors"`
	ShutDownCtxTimeout time.Duration `koanf:"shutdown_ctx_timeout"`
}

type Cors struct {
	AllowOrigins []string `koanf:"allow_origins"`
}

func New(config Config) Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.CORS.AllowOrigins,
	}))

	return Server{
		Router: e,
		Config: config,
	}
}

func (s Server) Start() error {
	return s.Router.Start(fmt.Sprintf(":%d", s.Config.Port))
}

func (s Server) Stop(ctx context.Context) error {
	return s.Router.Shutdown(ctx)
}
