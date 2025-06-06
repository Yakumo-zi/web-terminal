package sqlite

import (
	"context"
	"errors"

	"github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/ent/asset"
	"github.com/Yakumo-zi/web-terminal/ent/credential"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type credentialRepository struct {
	client *ent.CredentialClient
}

func newCredentialRepository(client *ent.Client) *credentialRepository {
	return &credentialRepository{
		client: client.Credential,
	}
}

func (r *credentialRepository) Create(ctx context.Context, cred *domain.Credential) error {
	if cred == nil {
		return errors.New("credential is nil")
	}
	_, err := r.client.Create().
		SetID(cred.Id).
		SetType(cred.Type).
		SetUsername(cred.Username).
		SetSecret(cred.Secret).
		SetAssetID(cred.Asset.Id).
		Save(ctx)
	return err
}

func (r *credentialRepository) Update(ctx context.Context, cred *domain.Credential) error {
	if cred == nil {
		return errors.New("credential is nil")
	}
	_, err := r.client.UpdateOneID(cred.Id).
		SetType(cred.Type).
		SetUsername(cred.Username).
		SetSecret(cred.Secret).
		Save(ctx)
	return err
}

func (r *credentialRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.client.DeleteOneID(id).Exec(ctx)
}

func (r *credentialRepository) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	_, err := r.client.Delete().Where(credential.IDIn(ids...)).Exec(ctx)
	return err
}

func (r *credentialRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Credential, error) {
	cred, err := r.client.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	ret := &domain.Credential{
		Id:        cred.ID,
		Type:      cred.Type,
		Username:  cred.Username,
		Secret:    cred.Secret,
		CreatedAt: cred.CreatedAt,
		UpdatedAt: cred.UpdatedAt,
	}
	return ret, nil
}

func (r *credentialRepository) List(ctx context.Context, options *repo.ListOptions) ([]*domain.Credential, int, error) {
	credentials, err := r.client.Query().Limit(options.Limit).Offset(options.Offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.client.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.Credential, len(credentials))
	for i, credential := range credentials {
		ret[i] = &domain.Credential{
			Id:        credential.ID,
			Type:      credential.Type,
			Secret:    credential.Secret,
			Username:  credential.Username,
			CreatedAt: credential.CreatedAt,
			UpdatedAt: credential.UpdatedAt,
		}
	}
	return ret, count, nil
}

func (c *credentialRepository) GetByAsset(ctx context.Context, assetID uuid.UUID, limit int, offset int) ([]*domain.Credential, int, error) {
	credentials, err := c.client.Query().Where(credential.HasAssetWith(asset.ID(assetID))).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := c.client.Query().Where(credential.HasAssetWith(asset.ID(assetID))).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.Credential, len(credentials))
	for i, credential := range credentials {
		ret[i] = &domain.Credential{
			Id:        credential.ID,
			Type:      credential.Type,
			Secret:    credential.Secret,
			Username:  credential.Username,
			CreatedAt: credential.CreatedAt,
			UpdatedAt: credential.UpdatedAt,
		}
	}
	return ret, count, nil
}
