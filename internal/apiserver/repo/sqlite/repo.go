package sqlite

import (
	"github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
)

type SqliteRepository struct {
	client *ent.Client
}

func NewSqliteRepository(client *ent.Client) *SqliteRepository {

	return &SqliteRepository{
		client: client,
	}
}
func (s SqliteRepository) Assets() repo.AssetRepository {
	return newAssetRepo(s.client.Asset)
}
func (s *SqliteRepository) Close() error {
	return s.client.Close()
}
