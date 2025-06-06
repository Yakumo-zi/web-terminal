package repo

import (
	"context"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/google/uuid"
)

type AssetGroupRepository interface {
	Create(context.Context, *domain.AssetGroup) error
	Update(context.Context, *domain.AssetGroup) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.AssetGroup, error)
	List(context.Context, *ListOptions) ([]*domain.AssetGroup, int, error)
	AddMembers(context.Context, uuid.UUID, []uuid.UUID) error
}
