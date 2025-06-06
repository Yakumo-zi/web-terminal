package service

import (
	"context"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type SessionService interface {
	Create(context.Context, *domain.Session) error
	Update(context.Context, *domain.Session) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.Session, error)
	List(context.Context, *ListOptions) ([]*domain.Session, int, error)
	GetByAsset(context.Context, uuid.UUID, int, int) ([]*domain.Session, int, error)
}

type sessionService struct {
	repo repo.SessionRepository
}

func newSessionService(repo repo.SessionRepository) *sessionService {
	return &sessionService{
		repo: repo,
	}
}

func (s *sessionService) Create(ctx context.Context, session *domain.Session) error {
	return s.repo.Create(ctx, session)
}

func (s *sessionService) Update(ctx context.Context, session *domain.Session) error {
	return s.repo.Update(ctx, session)
}

func (s *sessionService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *sessionService) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	return s.repo.DeleteCollection(ctx, ids)
}

func (s *sessionService) Get(ctx context.Context, id uuid.UUID) (*domain.Session, error) {
	return s.repo.Get(ctx, id)
}

func (s *sessionService) List(ctx context.Context, options *ListOptions) ([]*domain.Session, int, error) {
	return s.repo.List(ctx, &repo.ListOptions{Limit: options.Limit, Offset: options.Offset})
}

func (s *sessionService) GetByAsset(ctx context.Context, assetID uuid.UUID, limit int, offset int) ([]*domain.Session, int, error) {
	return s.repo.GetByAsset(ctx, assetID, limit, offset)
}
