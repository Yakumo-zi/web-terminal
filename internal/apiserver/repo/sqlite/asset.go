package sqlite

import (
	"context"
	"errors"

	"github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/ent/asset"
	"github.com/Yakumo-zi/web-terminal/ent/assetgroup"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type assetRepo struct {
	client *ent.AssetClient
}

func newAssetRepository(client *ent.Client) *assetRepo {
	return &assetRepo{client: client.Asset}
}

func (a assetRepo) Create(ctx context.Context, asset *domain.Asset) error {
	if asset == nil {
		return errors.New("asset is nil")
	}
	_, err := a.client.Create().
		SetID(asset.Id).
		SetName(asset.Name).
		SetType(asset.Type).
		SetIP(asset.Ip).
		SetPort(asset.Port).
		Save(ctx)
	return err
}

func (a assetRepo) Update(ctx context.Context, asset *domain.Asset) error {
	if asset == nil {
		return errors.New("asset is nil")
	}
	_, err := a.client.UpdateOneID(asset.Id).
		SetName(asset.Name).
		SetIP(asset.Ip).
		SetPort(asset.Port).
		SetType(asset.Type).
		Save(ctx)
	return err

}

func (a assetRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return a.client.DeleteOneID(id).Exec(ctx)
}

func (a assetRepo) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	_, err := a.client.Delete().Where(asset.IDIn(ids...)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a assetRepo) Get(ctx context.Context, id uuid.UUID) (*domain.Asset, error) {
	asset, err := a.client.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	ret := &domain.Asset{
		Ip:   asset.IP,
		Port: asset.Port,
		Name: asset.Name,
		Type: asset.Type,
		Id:   asset.ID,
	}
	return ret, nil
}

func (a assetRepo) List(ctx context.Context, options *repo.ListOptions) ([]*domain.Asset, int, error) {
	assets, err := a.client.Query().Limit(options.Limit).Offset(options.Offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := a.client.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.Asset, len(assets))
	for i, asset := range assets {
		ret[i] = &domain.Asset{
			Id:        asset.ID,
			Name:      asset.Name,
			Ip:        asset.IP,
			Port:      asset.Port,
			Type:      asset.Type,
			CreatedAt: asset.CreatedAt,
			UpdatedAt: asset.UpdatedAt,
		}
	}
	return ret, count, nil
}

func (a assetRepo) GetByGroup(ctx context.Context, groupID uuid.UUID, limit, offset int) ([]*domain.Asset, int, error) {
	assets, err := a.client.Query().Where(asset.HasGroupsWith(assetgroup.ID(groupID))).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := a.client.Query().Where(asset.HasGroupsWith(assetgroup.ID(groupID))).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.Asset, len(assets))
	for i, asset := range assets {
		ret[i] = &domain.Asset{
			Id:        asset.ID,
			Name:      asset.Name,
			Ip:        asset.IP,
			Port:      asset.Port,
			Type:      asset.Type,
			CreatedAt: asset.CreatedAt,
			UpdatedAt: asset.UpdatedAt,
		}
	}
	return ret, count, nil
}

func (a assetRepo) GetWithoutGroup(ctx context.Context, limit, offset int) ([]*domain.Asset, int, error) {
	assets, err := a.client.Query().Where(asset.Not(asset.HasGroups())).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := a.client.Query().Where(asset.Not(asset.HasGroups())).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.Asset, len(assets))
	for i, asset := range assets {
		ret[i] = &domain.Asset{
			Id:        asset.ID,
			Name:      asset.Name,
			Ip:        asset.IP,
			Port:      asset.Port,
			Type:      asset.Type,
			CreatedAt: asset.CreatedAt,
			UpdatedAt: asset.UpdatedAt,
		}
	}
	return ret, count, nil
}
