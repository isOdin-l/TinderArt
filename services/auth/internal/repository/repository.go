package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	builder "github.com/isOdin-l/TinderArt/services/auth/internal/database/sqlc"
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound = errors.New("not found")
)

type AuthRepository struct {
	q *builder.Queries
}

func NewRepository(db builder.DBTX) *AuthRepository {
	return &AuthRepository{q: builder.New(db)}
}

func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (*entities.Login, error) {
	row, err := r.q.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entities.Login{
		UserId:   FromPgUUID(&row.ID),
		Password: row.Password,
	}, nil
}

func (r *AuthRepository) SaveRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string) error {
	return r.q.SaveRefreshToken(ctx,
		builder.SaveRefreshTokenParams{
			ID:           ToPgUUID(&userId),
			RefreshToken: refreshToken,
		},
	)
}
