package psql

import (
	"context"
	"fmt"
	"time"

	l "github.com/MohammadBohluli/social-content-app/adapter/logger"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Driver         string `koanf:"postgres"`
	Username       string `koanf:"username"`
	Password       string `koanf:"password"`
	Host           string `koanf:"host"`
	Port           int    `koanf:"port"`
	DBName         string `koanf:"db_name"`
	SSLMode        string `koanf:"ssl_mode"`
	ConnectTimeout int    `koanf:"connection_timeout"`
}

type DB struct {
	config Config
	conn   *pgx.Conn
}

func (p *DB) Conn() *pgx.Conn {
	return p.conn
}

func New(config Config, logger l.Logger) *DB {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		logger.Fatal("❌ can't connect to postgreSQL", "error", err)
	}

	conn.Config().ConnectTimeout = time.Duration(config.ConnectTimeout) * time.Second

	logger.Info("✅ postgreSQL up and running")

	return &DB{config: config, conn: conn}
}
