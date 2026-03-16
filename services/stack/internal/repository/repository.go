package repository

import (
	"context"

	"github.com/isOdin-l/TinderArt/services/stack/internal/database/sqlc"
	"github.com/isOdin-l/TinderArt/services/stack/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q *sqlc.Queries
}

func NewRepository(db sqlc.DBTX) *Repository {
	return &Repository{q: sqlc.New(db)}
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
