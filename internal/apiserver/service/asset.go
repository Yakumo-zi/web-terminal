package service

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type AssetService interface {
	Create(context.Context, *domain.Asset) error
	Update(context.Context, *domain.Asset) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.Asset, error)
	List(context.Context, *ListOptions) ([]*domain.Asset, int, error)
}

type assetService struct {
	repo repo.AssetRepository
}

func newAssetService(repo repo.AssetRepository) *assetService {
	return &assetService{
		repo: repo,
	}
}

func (a *assetService) Create(ctx context.Context, asset *domain.Asset) error {
	return a.repo.Create(ctx, asset)
}

func (a *assetService) Update(ctx context.Context, asset *domain.Asset) error {
	return a.repo.Update(ctx, asset)
}

func (a *assetService) Delete(ctx context.Context, id uuid.UUID) error {
	return a.repo.Delete(ctx, id)
}

func (a *assetService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	return a.repo.DeleteCollection(ctx, ids)
}

func (a *assetService) Get(ctx context.Context, s uuid.UUID) (*domain.Asset, error) {
	return a.repo.Get(ctx, s)
}

func (a *assetService) List(ctx context.Context, options *ListOptions) ([]*domain.Asset, int, error) {
	return a.repo.List(ctx, &repo.ListOptions{Limit: options.Limit, Offset: options.Offset})
}
