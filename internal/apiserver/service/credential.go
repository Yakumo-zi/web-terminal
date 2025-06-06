package service

import (
	"context"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type CredentialService interface {
	Create(context.Context, *domain.Credential) error
	Update(context.Context, *domain.Credential) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.Credential, error)
	List(context.Context, *ListOptions) ([]*domain.Credential, int, error)
	GetByAsset(context.Context, uuid.UUID, int, int) ([]*domain.Credential, int, error)
}

type credentialService struct {
	repo repo.CredentialRepository
}

func newCredentialService(repo repo.CredentialRepository) *credentialService {
	return &credentialService{
		repo: repo,
	}
}

func (c *credentialService) Create(ctx context.Context, credential *domain.Credential) error {
	return c.repo.Create(ctx, credential)
}

func (c *credentialService) Update(ctx context.Context, credential *domain.Credential) error {
	return c.repo.Update(ctx, credential)
}

func (c *credentialService) Delete(ctx context.Context, id uuid.UUID) error {
	return c.repo.Delete(ctx, id)
}

func (c *credentialService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	return c.repo.DeleteCollection(ctx, ids)
}

func (c *credentialService) Get(ctx context.Context, id uuid.UUID) (*domain.Credential, error) {
	return c.repo.Get(ctx, id)
}

func (c *credentialService) List(ctx context.Context, options *ListOptions) ([]*domain.Credential, int, error) {
	return c.repo.List(ctx, &repo.ListOptions{Limit: options.Limit, Offset: options.Offset})
}

func (c *credentialService) GetByAsset(ctx context.Context, assetID uuid.UUID, limit int, offset int) ([]*domain.Credential, int, error) {
	return c.repo.GetByAsset(ctx, assetID, limit, offset)
}
