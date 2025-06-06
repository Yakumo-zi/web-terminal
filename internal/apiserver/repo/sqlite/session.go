package sqlite

import (
	"context"
	"errors"
	"time"

	"github.com/Yakumo-zi/web-terminal/ent"
	"github.com/Yakumo-zi/web-terminal/ent/asset"
	"github.com/Yakumo-zi/web-terminal/ent/session"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/repo"
	"github.com/google/uuid"
)

type sessionRepository struct {
	client *ent.SessionClient
}

func newSessionRepository(client *ent.Client) *sessionRepository {
	return &sessionRepository{
		client: client.Session,
	}
}

func (r *sessionRepository) Create(ctx context.Context, sess *domain.Session) error {
	if sess == nil {
		return errors.New("session is nil")
	}
	_, err := r.client.Create().
		SetID(sess.Id).
		SetType(sess.Type).
		SetStatus(sess.Status).
		SetCreatedAt(sess.CreatedAt).
		SetUpdatedAt(sess.UpdatedAt).
		SetAssetID(sess.Asset.Id).
		SetCredentialID(sess.Credential.Id).
		SetStopedAt(time.Time{}).
		Save(ctx)
	return err
}

func (r *sessionRepository) Update(ctx context.Context, sess *domain.Session) error {
	if sess == nil {
		return errors.New("session is nil")
	}
	_, err := r.client.UpdateOneID(sess.Id).
		SetType(sess.Type).
		SetStatus(sess.Status).
		SetUpdatedAt(sess.UpdatedAt).
		SetStopedAt(sess.StopedAt).
		Save(ctx)
	return err
}

func (r *sessionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.client.DeleteOneID(id).Exec(ctx)
}

func (r *sessionRepository) DeleteCollection(ctx context.Context, ids []uuid.UUID) error {
	_, err := r.client.Delete().Where(session.IDIn(ids...)).Exec(ctx)
	return err
}

func (r *sessionRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Session, error) {
	sess, err := r.client.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	var stopedAt time.Time
	if sess.StopedAt != nil {
		stopedAt = *sess.StopedAt
	}
	ret := &domain.Session{
		Id:        sess.ID,
		Type:      sess.Type,
		Status:    sess.Status,
		CreatedAt: sess.CreatedAt,
		UpdatedAt: sess.UpdatedAt,
		StopedAt:  stopedAt,
	}
	return ret, nil
}

func (r *sessionRepository) List(ctx context.Context, options *repo.ListOptions) ([]*domain.Session, int, error) {
	sessions, err := r.client.Query().Limit(options.Limit).Offset(options.Offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.client.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.Session, len(sessions))
	for i, sess := range sessions {
		var stopedAt time.Time
		if sess.StopedAt != nil {
			stopedAt = *sess.StopedAt
		}
		ret[i] = &domain.Session{
			Id:        sess.ID,
			Type:      sess.Type,
			Status:    sess.Status,
			CreatedAt: sess.CreatedAt,
			UpdatedAt: sess.UpdatedAt,
			StopedAt:  stopedAt,
		}
	}
	return ret, count, nil
}

func (r *sessionRepository) GetByAsset(ctx context.Context, assetID uuid.UUID, limit int, offset int) ([]*domain.Session, int, error) {
	sessions, err := r.client.Query().Where(session.HasAssetWith(asset.ID(assetID))).Limit(limit).Offset(offset).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.client.Query().Where(session.HasAssetWith(asset.ID(assetID))).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ret := make([]*domain.Session, len(sessions))
	for i, sess := range sessions {
		var stopedAt time.Time
		if sess.StopedAt != nil {
			stopedAt = *sess.StopedAt
		}
		ret[i] = &domain.Session{
			Id:        sess.ID,
			Type:      sess.Type,
			Status:    sess.Status,
			CreatedAt: sess.CreatedAt,
			UpdatedAt: sess.UpdatedAt,
			StopedAt:  stopedAt,
		}
	}
	return ret, count, nil
}
