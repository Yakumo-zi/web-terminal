package service

import (
	"context"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type AssetGroupService interface {
	Create(context.Context, *domain.AssetGroup) error
	Update(context.Context, *domain.AssetGroup) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.AssetGroup, error)
	List(context.Context, *ListOptions) ([]*domain.AssetGroup, int, error)
	AddMembers(context.Context, uuid.UUID, []uuid.UUID) error
}

type assetGroupService struct {
	repo repo.AssetGroupRepository
}

func newAssetGroupService(repo repo.AssetGroupRepository) *assetGroupService {
	return &assetGroupService{
		repo: repo,
	}
}

func (a *assetGroupService) Create(ctx context.Context, group *domain.AssetGroup) error {
	return a.repo.Create(ctx, group)
}

func (a *assetGroupService) Update(ctx context.Context, group *domain.AssetGroup) error {
	return a.repo.Update(ctx, group)
}

func (a *assetGroupService) Delete(ctx context.Context, id uuid.UUID) error {
	return a.repo.Delete(ctx, id)
}

func (a *assetGroupService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	return a.repo.DeleteCollection(ctx, ids)
}

func (a *assetGroupService) Get(ctx context.Context, id uuid.UUID) (*domain.AssetGroup, error) {
	return a.repo.Get(ctx, id)
}

func (a *assetGroupService) List(ctx context.Context, options *ListOptions) ([]*domain.AssetGroup, int, error) {
	return a.repo.List(ctx, &repo.ListOptions{Limit: options.Limit, Offset: options.Offset})
}

func (a *assetGroupService) AddMembers(ctx context.Context, groupID uuid.UUID, memberIDs []uuid.UUID) error {
	return a.repo.AddMembers(ctx, groupID, memberIDs)
}
