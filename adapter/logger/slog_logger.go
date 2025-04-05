package logger

import (
	"log/slog"
	"os"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() *SlogLogger {

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &SlogLogger{
		logger: slog.New(handler),
	}
}

func (s *SlogLogger) Info(msg string, keysAndValues ...any) {
	s.logger.Info(msg, keysAndValues...)
}

func (s *SlogLogger) Error(msg string, keysAndValues ...any) {
	s.logger.Error(msg, keysAndValues...)
}

func (s *SlogLogger) Fatal(msg string, keysAndValues ...any) {
	s.logger.Error(msg, keysAndValues...)
	os.Exit(1)
}
