package service

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	AssetService AssetService
}

func NewService() *Service {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	repo := sqlite.NewSqliteRepository(client)

	aSvc := newAssetService(repo.Assets())
	svc := &Service{
		AssetService: aSvc,
	}
	return svc
}
