package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/isOdin-l/TinderArt/services/profile/internal/repository/mappers"
	"github.com/jackc/pgx/v5/pgtype"
)

type QueryBuilder interface {
	// Profile
	CreateProfile(ctx context.Context, arg db_models.CreateProfileParams) error
	GetProfile(ctx context.Context, id pgtype.UUID) (db_models.GetProfileRow, error) // TODO: should be with photos urls
	UpdateProfile(ctx context.Context, arg db_models.UpdateProfileParams) (db_models.UpdateProfileRow, error)
	DeleteProfile(ctx context.Context, id pgtype.UUID) error

	// Preferences
	CreatePreferences(ctx context.Context, arg db_models.CreatePreferencesParams) error
	UpdatePreferences(ctx context.Context, arg db_models.UpdatePreferencesParams) (db_models.Preference, error)

	// Photos
	CreatePhotos(ctx context.Context, arg db_models.CreatePhotosParams) error
	GetPhotos(ctx context.Context, profileID pgtype.UUID) ([]pgtype.UUID, error)
	DeletePhotos(ctx context.Context, arg db_models.DeletePhotosParams) error

	// FavArtStyles
	CreateFavArtStyle(ctx context.Context, arg db_models.CreateFavArtStyleParams) error
}

type Repository struct {
	query QueryBuilder
}

func NewRepository(db db_models.DBTX) *Repository {
	return &Repository{query: db_models.New(db)}
}

// PROFILE
func (repo *Repository) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	return repo.query.CreateProfile(ctx, *mappers.FromEntityToCreateProfile(profile))
}

func (repo *Repository) GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	row, errDb := repo.query.GetProfile(ctx, db_models.ToPgUUID(&userId))
	if errDb != nil {
		return nil, errDb
	}
	return mappers.FromGetProfileToEntity(&row), nil
}

func (repo *Repository) UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error) {
	row, errDb := repo.query.UpdateProfile(ctx, *mappers.FromEntityToUpdateProfileParams(profile))
	if errDb != nil {
		return nil, errDb
	}
	return mappers.FromUpdateProfileToEntity(&row), nil
}

func (repo *Repository) DeleteProfile(ctx context.Context, userId uuid.UUID) error {
	return repo.query.DeleteProfile(ctx, db_models.ToPgUUID(&userId))
}

// PREFERENCES
func (repo *Repository) CreatePreferences(ctx context.Context, profile *entities.Profile) error {
	return repo.query.CreatePreferences(ctx, *mappers.FromEntityToCreatePrefParams(profile))
}

func (repo *Repository) UpdatePreferences(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error) {
	db, errDb := repo.query.UpdatePreferences(ctx, *mappers.FromEntityToUpdatePref(profile))
	if errDb != nil {
		return nil, errDb
	}
	return mappers.FromPreferenceToEntity(&db), nil
}

// Favourite art styles
func (repo *Repository) CreateFavArtStyle(ctx context.Context, profile *entities.Profile) error {
	return repo.query.CreateFavArtStyle(ctx, *mappers.FromEntityToCreateFavArtStyles(profile))
}

// PHOTOS
func (repo *Repository) CreatePhotos(ctx context.Context, profile *entities.Profile) error {
	return repo.query.CreatePhotos(ctx, *mappers.FromEntityToCreatePhotos(profile))
}

func (repo *Repository) GetPhotos(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	photosUrls, errDb := repo.query.GetPhotos(ctx, db_models.ToPgUUID(&userId))
	return mappers.FromGetPhotosToEntity(&photosUrls), errDb
}

// func (repo *Repository) DeletePhotos(ctx context.Context, profile *entities.Profile) error {
// 	return repo.query.DeletePhotos(ctx, *FromEntityToDeletePhotos(profile))
// }
