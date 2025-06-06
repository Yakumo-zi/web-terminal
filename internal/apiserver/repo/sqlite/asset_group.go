package sqlite

import (
	"context"
	"errors"

	"github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/ent/assetgroup"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type assetGroupRepository struct {
	client      *ent.Client
	groupClient *ent.AssetGroupClient
}

func newAssetGroupRepository(client *ent.Client) *assetGroupRepository {
	return &assetGroupRepository{
		client:      client,
		groupClient: client.AssetGroup,
	}
}

func (r *assetGroupRepository) Create(ctx context.Context, group *domain.AssetGroup) error {
	if group == nil {
		return errors.New("group is nil")
	}
	_, err := r.groupClient.Create().
		SetID(group.Id).
		SetName(group.Name).
		SetCreatedAt(group.CreatedAt).
		SetUpdatedAt(group.UpdatedAt).
		Save(ctx)
	return err
}

func (r *assetGroupRepository) Update(ctx context.Context, group *domain.AssetGroup) error {
	if group == nil {
		return errors.New("group is nil")
	}
	_, err := r.groupClient.UpdateOneID(group.Id).
		SetName(group.Name).
		SetUpdatedAt(group.UpdatedAt).
		Save(ctx)
	return err
}

func (r *assetGroupRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.groupClient.DeleteOneID(id).Exec(ctx)
}

func (r *assetGroupRepository) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	_, err := r.groupClient.Delete().Where(assetgroup.IDIn(ids...)).Exec(ctx)
	return err
}

func (r *assetGroupRepository) Get(ctx context.Context, id uuid.UUID) (*domain.AssetGroup, error) {
	group, err := r.groupClient.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	ret := &domain.AssetGroup{
		Id:        group.ID,
		Name:      group.Name,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
	}
	return ret, nil
}

func (r *assetGroupRepository) List(ctx context.Context, options *repo.ListOptions) ([]*domain.AssetGroup, int, error) {
	groups, err := r.groupClient.Query().Limit(options.Limit).Offset(options.Offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.groupClient.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.AssetGroup, len(groups))
	for i, group := range groups {
		ret[i] = &domain.AssetGroup{
			Id:        group.ID,
			Name:      group.Name,
			CreatedAt: group.CreatedAt,
			UpdatedAt: group.UpdatedAt,
		}
	}
	return ret, count, nil
}

func (a *assetGroupRepository) AddMembers(ctx context.Context, groupID uuid.UUID, memberIDs []uuid.UUID) error {
	_, err := a.groupClient.UpdateOneID(groupID).AddAssetIDs(memberIDs...).Save(ctx)
	return err
}
