package service

import (
	"context"
	"encoding/json"
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
	LRange(ctx context.Context, key string, start, end int64) ([]any, error)
}

type IRepository interface {
	// Profile
	CreateProfile(ctx context.Context, profile *entities.Profile) error
	GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)
	UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error)
	DeleteProfile(ctx context.Context, userId uuid.UUID) error

	// Preference
	CreatePreferences(ctx context.Context, profile *entities.Profile) error

	// Photos
	CreatePhotos(ctx context.Context, profile *entities.Profile) error

	// FavArtStyles
	CreateFavArtStyle(ctx context.Context, profile *entities.Profile) error
}

type IStorage interface {
	PutObject(ctx context.Context, bucket, key *string, body io.Reader) error
	GetObject(ctx context.Context, bucket, key *string) ([]byte, error)
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
	for idx := range len(profile.FavArtStyles) {
		profile.FavArtStylesIds[idx], errPrep = uuid.NewV7()
		if errPrep != nil {
			return errPrep
		}
	}

	// Create photos data
	for idx := range len(profile.PhotoFiles) {
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
	result, errCall := s.grpc_client.CreateUser(ctx, &grpc_auth.CreateUserRequest{
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
	cacheKey := fmt.Sprintf("profile:%s", userId.String())

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

	return s.repo.UpdateProfile(ctx, profile)
}
func (s *Service) DeleteProfile(ctx context.Context, userId uuid.UUID) error {
	profile, errGet := s.repo.GetProfile(ctx, userId)
	if errGet != nil {
		return errGet
	}

	// + удалять фото из s3

	return s.repo.DeleteProfile(ctx, profile.UserId)
}

func (s *Service) genPasswordHash(password string) (string, error) {
	hash, errHash := bcrypt.GenerateFromPassword([]byte(password), s.cfg.HashMinCost)
	return string(hash), errHash
}

// -- Get stack --
// For future:
// 	// 0 - start point; -1 - like the last element, len(arg)-1
// res, errCache := s.cache.LRange(ctx, userId.String(), 0, -1)
// if errCache == nil{
// 	return res
// }
