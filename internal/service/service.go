package service

import (
	"github.com/Yakumo-zi/web-terminal/pkg/logger"
	"log/slog"
	"os"
)

type Service struct {
	BaseLogger *slog.Logger
	WebLogger  *slog.Logger
}

func NewService() *Service {
	log := logger.NewLogger(os.Stdout)
	webLogger := logger.NewWebLogger(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	return &Service{BaseLogger: log, WebLogger: webLogger}
}
