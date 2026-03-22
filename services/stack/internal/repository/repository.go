package repository

import (
	"context"

	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/stack/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

type QueryBuilder interface {
	FindMatches(ctx context.Context, id pgtype.UUID) ([]pgtype.UUID, error)
	GetAllProfiles(ctx context.Context) ([]pgtype.UUID, error)
}

type Repository struct {
	q QueryBuilder
}

func NewRepository(db db_models.DBTX) *Repository {
	return &Repository{q: db_models.New(db)}
}

func (r *Repository) GetAllProfiles(ctx context.Context) ([]pgtype.UUID, error) {
	return r.q.GetAllProfiles(ctx)
}

func (r *Repository) FindMatches(ctx context.Context, id pgtype.UUID) ([]interface{}, error) {
	matchesDb, errDb := r.q.FindMatches(ctx, id)
	if errDb != nil {
		return nil, errDb
	}

	return entities.FromPgTypeToInterface(matchesDb), nil
}
