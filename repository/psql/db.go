package psql

import (
	"database/sql"
	"fmt"
	"os"

	l "github.com/MohammadBohluli/social-content-app/adapter/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
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
	conn   *sql.DB
}

func (db *DB) Conn() *sql.DB {
	return db.conn
}

func (db *DB) Close() error {
	return db.conn.Close()
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

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Fatal("❌ can't connect to postgreSQL", "error", err)
		os.Exit(1)
	}

	logger.Info("✅ postgreSQL up and running")

	return &DB{config: config, conn: db}
}
