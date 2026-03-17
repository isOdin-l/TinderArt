package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/services/profile/internal/database/sqlc"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
)

type Repository struct {
	q *sqlc.Queries
}

func NewRepository(db sqlc.DBTX) *Repository {
	return &Repository{q: sqlc.New(db)}
}

func (r *Repository) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	return r.q.CreateProfile(ctx, *FromEntityToCreateProfile(profile))
}

func (r *Repository) GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	row, errDb := r.q.GetProfile(ctx, ToPgUUID(&userId))
	if errDb != nil {
		return nil, errDb
	}
	return FromGetProfileRowToEntity(&row), nil
}

func (r *Repository) UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error) {
	row, errDb := r.q.UpdateProfile(ctx, *FromEntityToUpdateProfileParams(profile))
	if errDb != nil {
		return nil, errDb
	}
	return FromUpdateResultsToEntity(&row), nil
}
func (r *Repository) DeleteProfile(ctx context.Context, userId uuid.UUID) error {
	return r.q.DeleteProfile(ctx, ToPgUUID(&userId))
}
