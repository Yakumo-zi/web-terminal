package sqlite

import (
	"github.com/Yakumo-zi/web-terminal/ent"
	_ "github.com/Yakumo-zi/web-terminal/ent/runtime"
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

func (r *SqliteRepository) Assets() repo.AssetRepository {
	return newAssetRepository(r.client)
}

func (r *SqliteRepository) AssetGroups() repo.AssetGroupRepository {
	return newAssetGroupRepository(r.client)
}

func (r *SqliteRepository) Credentials() repo.CredentialRepository {
	return newCredentialRepository(r.client)
}

func (r *SqliteRepository) Sessions() repo.SessionRepository {
	return newSessionRepository(r.client)
}

func (s *SqliteRepository) Close() error {
	return s.client.Close()
}
