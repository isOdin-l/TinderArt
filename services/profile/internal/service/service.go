package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	"github.com/isOdin-l/TinderArt/services/profile/config"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type ICache interface {
	Set(ctx context.Context, key string, value any, timeExpire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	RPush(ctx context.Context, key string, args ...any) error
	LPopCount(ctx context.Context, key string, count int) ([]string, error)
	Del(ctx context.Context, keys ...string) error
}

type IRepository interface {
	// Profile
	CreateProfile(ctx context.Context, profile *entities.Profile) error
	GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)
	UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error)
	DeleteProfile(ctx context.Context, userId uuid.UUID) error

	// Preference
	CreatePreferences(ctx context.Context, profile *entities.Profile) error
	UpdatePreferences(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error)

	// Photos
	CreatePhotos(ctx context.Context, profile *entities.Profile) error
	GetPhotos(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)

	// FavArtStyles
	CreateFavArtStyle(ctx context.Context, profile *entities.Profile) error

	//Stack
	GetStack(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)
}

type IStorage interface {
	PutObject(ctx context.Context, bucket, key *string, body io.Reader) error
	DeleteObjects(ctx context.Context, bucket *string, keys *[]string) error
}

type ITxManagaer interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error)
}

type Service struct {
	repo        IRepository
	storage     IStorage
	grpc_client grpc_auth.AuthServiceClient
	cfg         *config.Config
	txMng       ITxManagaer
	cache       ICache
}

func NewService(cfg *config.Config, repo IRepository, storage IStorage, grpc_client grpc_auth.AuthServiceClient, txMng ITxManagaer, cache ICache) *Service {
	return &Service{
		repo:        repo,
		storage:     storage,
		grpc_client: grpc_client,
		cfg:         cfg,
		txMng:       txMng,
		cache:       cache,
	}
}

func (s *Service) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	// Password hashing
	var errPrep error
	profile.Password, errPrep = s.genPasswordHash(profile.Password)
	if errPrep != nil {
		return errPrep
	}

	// Create new userID
	profile.UserId, errPrep = uuid.NewV7()
	if errPrep != nil {
		return errPrep
	}

	// Create uuid for favart
	for idx := range profile.FavArtStyles {
		profile.FavArtStylesIds[idx], errPrep = uuid.NewV7()
		if errPrep != nil {
			return errPrep
		}
	}

	// Create photos data
	for idx := range profile.PhotoFiles {
		profile.PhotosIds[idx], errPrep = uuid.NewV7()
		if errPrep != nil {
			return errPrep
		}

		profile.PhotoUrls[idx] = s.cfg.ConfigRustFS.ObjPath(profile.PhotosIds[idx].String())
	}

	_, errTx := s.txMng.WithTx(ctx, func(ctx context.Context) (any, error) {
		//Call database to insert data into profile, preferences, photos, fav_art_styles
		if errCreate := s.repo.CreateProfile(ctx, profile); errCreate != nil {
			return nil, errCreate
		}
		if errCreate := s.repo.CreatePreferences(ctx, profile); errCreate != nil {
			return nil, errCreate
		}
		if errCreate := s.repo.CreatePhotos(ctx, profile); errCreate != nil {
			return nil, errCreate
		}
		if errCreate := s.repo.CreateFavArtStyle(ctx, profile); errCreate != nil {
			return nil, errCreate
		}

		// Put photos to s3
		var filename string
		for idx, fileHeader := range profile.PhotoFiles {
			file, errOpen := fileHeader.Open()
			if errOpen != nil {
				return nil, errOpen
			}
			defer file.Close()

			filename = profile.PhotosIds[idx].String()

			if errPut := s.storage.PutObject(ctx, &s.cfg.RustFSBucketName, &filename, file); errPut != nil {
				return nil, errPut
			}
		}

		return nil, nil
	})
	if errTx != nil {
		return errTx
	}

	// Call Auth by gRPC to sign refresh and access tokens
	result, errCall := s.grpc_client.SignTokens(ctx, &grpc_auth.RequestSignTokens{
		UserId: profile.UserId.String(),
	})
	if errCall != nil {
		return errCall
	}

	// Updating entity fields
	profile.AccessToken = result.AccessToken
	profile.RefreshToken = result.RefreshToken

	return nil
}
func (s *Service) GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	cacheKey := fmt.Sprintf("profile:%s", userId)

	// Try to get user from redis
	userCache, errSearch := s.cache.Get(ctx, cacheKey)
	if errSearch == nil {
		var user entities.Profile
		return &user, json.Unmarshal([]byte(userCache), &user)
	}

	// Get user from DB
	userDb, errDb := s.repo.GetProfile(ctx, userId)
	if errDb != nil {
		return nil, errDb
	}
	data, _ := json.Marshal(userDb)

	// Set user to cache and return data from DB
	return userDb, s.cache.Set(ctx, cacheKey, data, 1*time.Hour)
}
func (s *Service) UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error) {
	if profile.Password != nil {
		tmpPass, errHash := s.genPasswordHash(*profile.Password)
		if errHash != nil {
			return nil, errHash
		}
		profile.Password = &tmpPass
	}

	res, errTx := s.txMng.WithTx(ctx, func(ctx context.Context) (any, error) {
		_, errUpdate := s.repo.UpdateProfile(ctx, profile)
		if errUpdate != nil {
			return nil, errUpdate
		}

		_, errUpdate = s.repo.UpdatePreferences(ctx, profile)
		if errUpdate != nil {
			return nil, errUpdate
		}

		return s.repo.GetProfile(ctx, profile.UserId)
	})
	if errTx != nil {
		return nil, errTx
	}

	resDb := res.(*entities.Profile)

	// Prepare data for writing to redis
	profileKey := fmt.Sprintf("profile:%s", profile.UserId.String())
	data, _ := json.Marshal(resDb)

	return resDb, s.cache.Set(ctx, profileKey, data, 1*time.Hour)
}
func (s *Service) DeleteProfile(ctx context.Context, userId uuid.UUID) error {
	stack_key := fmt.Sprintf("stack:%s", userId)
	profile_key := fmt.Sprintf("profile:%s", userId)

	if errCache := s.cache.Del(ctx, stack_key, profile_key); errCache != nil {
		return errCache
	}

	_, errTx := s.txMng.WithTx(ctx, func(ctx context.Context) (any, error) {
		photos, errPhotos := s.repo.GetPhotos(ctx, userId)
		if errPhotos != nil {
			return nil, errPhotos
		}

		if errDb := s.repo.DeleteProfile(ctx, userId); errDb != nil {
			return nil, errDb
		}

		// Preparation before deletion in storage
		photosIds := make([]string, len(photos.PhotosIds))
		for i := range len(photos.PhotosIds) {
			photosIds[i] = photos.PhotosIds[i].String()
		}

		if errStor := s.storage.DeleteObjects(ctx, &s.cfg.RustFSBucketName, &photosIds); errStor != nil {
			return nil, errStor
		}

		return nil, nil
	})

	return errTx
}

func (s *Service) GetStack(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	cacheKey := fmt.Sprintf("stack:%s", userId)

	// Try to get stack from cache
	matchesCache, errCache := s.cache.LPopCount(ctx, cacheKey, 20)
	if errCache == nil {
		return &entities.Profile{Matches: matchesCache}, nil
	}

	entity, errRepo := s.repo.GetStack(ctx, userId)
	if errRepo != nil {
		return nil, errRepo
	}
	if len(entity.Matches) != 0 {
		return entity, s.cache.RPush(ctx, cacheKey, entity.Matches)
	}
	return entity, errors.New("no rows in result set")
}

func (s *Service) genPasswordHash(password string) (string, error) {
	hash, errHash := bcrypt.GenerateFromPassword([]byte(password), s.cfg.HashMinCost)
	return string(hash), errHash
}
