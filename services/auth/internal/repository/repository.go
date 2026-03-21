package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrNotFound = errors.New("not found")
)

type QueryBuilder interface {
	GetUserByUsername(ctx context.Context, username pgtype.Text) (db_models.GetUserByUsernameRow, error)
	SaveRefreshToken(ctx context.Context, arg db_models.SaveRefreshTokenParams) error
}

type AuthRepository struct {
	query QueryBuilder
}

func NewRepository(db db_models.DBTX) *AuthRepository {
	return &AuthRepository{query: db_models.New(db)}
}

func (repo *AuthRepository) GetUserByUsername(ctx context.Context, username string) (*entities.Login, error) {
	row, err := repo.query.GetUserByUsername(ctx, ToPgText(&username))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entities.Login{
		UserId:   FromPgUUID(&row.ID),
		Password: FromPgText(&row.Password),
	}, nil
}

func (repo *AuthRepository) SaveRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string) error {
	return repo.query.SaveRefreshToken(ctx,
		db_models.SaveRefreshTokenParams{
			ID:           ToPgUUID(&userId),
			RefreshToken: ToPgText(&refreshToken),
		},
	)
}
