package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

type QueryBuilder interface {
	// Create
	CreatePhotos(ctx context.Context, arg db_models.CreatePhotosParams) error
	CreatePreferences(ctx context.Context, arg db_models.CreatePreferencesParams) error
	CreateProfile(ctx context.Context, arg db_models.CreateProfileParams) error

	// Get
	GetProfile(ctx context.Context, id pgtype.UUID) (db_models.GetProfileRow, error)

	// Update
	UpdatePreferences(ctx context.Context, arg db_models.UpdatePreferencesParams) (db_models.Preference, error)
	UpdateProfile(ctx context.Context, arg db_models.UpdateProfileParams) (db_models.UpdateProfileRow, error)

	// Delete
	DeletePhotos(ctx context.Context, arg db_models.DeletePhotosParams) error
	DeleteProfile(ctx context.Context, id pgtype.UUID) error
}

type Repository struct {
	query QueryBuilder
}

func NewRepository(db db_models.DBTX) *Repository {
	return &Repository{query: db_models.New(db)}
}

func (repo *Repository) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	return repo.query.CreateProfile(ctx, *FromEntityToCreateProfile(profile))
}

func (repo *Repository) GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	row, errDb := repo.query.GetProfile(ctx, ToPgUUID(&userId))
	if errDb != nil {
		return nil, errDb
	}
	return FromGetProfileRowToEntity(&row), nil
}

func (repo *Repository) UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error) {
	row, errDb := repo.query.UpdateProfile(ctx, *FromEntityToUpdateProfileParams(profile))
	if errDb != nil {
		return nil, errDb
	}
	return FromUpdateResultsToEntity(&row), nil
}
func (repo *Repository) DeleteProfile(ctx context.Context, userId uuid.UUID) error {
	return repo.query.DeleteProfile(ctx, ToPgUUID(&userId))
}

// CRUD preferences, CRUD photo
