package repo

import (
	"context"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/google/uuid"
)

type SessionRepository interface {
	Create(context.Context, *domain.Session) error
	Update(context.Context, *domain.Session) error
	Delete(context.Context, uuid.UUID) error
	DeleteCollection(context.Context, []uuid.UUID) error
	Get(context.Context, uuid.UUID) (*domain.Session, error)
	List(context.Context, *ListOptions) ([]*domain.Session, int, error)
	GetByAsset(context.Context, uuid.UUID, int, int) ([]*domain.Session, int, error)
}
