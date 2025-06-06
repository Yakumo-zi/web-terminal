package repo

import (
	"context"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/google/uuid"
)

type CredentialRepository interface {
	Create(context.Context, *domain.Credential) error
	Update(context.Context, *domain.Credential) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.Credential, error)
	List(context.Context, *ListOptions) ([]*domain.Credential, int, error)
	GetByAsset(context.Context, uuid.UUID, int, int) ([]*domain.Credential, int, error)
}
