package service

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
)

type AssetService interface {
	Create(context.Context, *domain.Asset) error
	Update(context.Context, *domain.Asset) error
	Delete(context.Context, string) error
	DeleteCollection(context.Context, []string) error
	Get(context.Context, string) (*domain.Asset, error)
	List(context.Context, *ListOptions) ([]*domain.Asset, int, error)
}

type assetService struct {
	repo repo.Repository
}
