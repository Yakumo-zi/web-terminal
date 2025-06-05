package service

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo/sqlite"
	"github.com/Yakumo-zi/web-terminal/pkg/logger"
	"log/slog"
	"os"
)

type Service struct {
	BaseLogger *slog.Logger
	WebLogger  *slog.Logger
	Repo       repo.Repository
}

func NewService() *Service {
	svc := &Service{}
	initBaseLog(svc)
	initWebLog(svc)
	initRepo(svc)
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

func initRepo(svc *Service) {
	client, err := ent.Open("sqlite3", "file:ent?_journal=PERSIST&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	repo := sqlite.NewSqliteRepository(client)
	svc.Repo = repo
}
