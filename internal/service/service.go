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
	svc := &Service{}
	initBaseLog(svc)
	initWebLog(svc)
	return svc
}

func initBaseLog(svc *Service) {
	log := logger.NewLogger(os.Stdout)
	svc.BaseLogger = log
}

func initWebLog(svc *Service) {
	log := logger.NewWebLogger(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	svc.WebLogger = log
}
