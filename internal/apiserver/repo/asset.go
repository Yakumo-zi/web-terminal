package repo

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/google/uuid"
)

type AssetRepository interface {
	Create(context.Context, *domain.Asset) error
	Update(context.Context, *domain.Asset) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.Asset, error)
	List(context.Context, *ListOptions) ([]*domain.Asset, int, error)
}
