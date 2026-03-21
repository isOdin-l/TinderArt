package service

import (
	"context"
	"io"

	"github.com/google/uuid"
	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	"github.com/isOdin-l/TinderArt/services/profile/config"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type ICache interface {
	LRange(ctx context.Context, key string, start, end int64) ([]any, error)
}

type IRepository interface {
	CreateProfile(ctx context.Context, profile *entities.Profile) error
	GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)
	UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error)
	DeleteProfile(ctx context.Context, userId uuid.UUID) error
}

type IStorage interface {
	PutObject(ctx context.Context, bucket, key *string, body io.Reader) error
	GetObject(ctx context.Context, bucket, key *string) ([]byte, error)
}

type Service struct {
	repo        IRepository
	storage     IStorage
	grpc_client grpc_auth.AuthServiceClient
	cfg         *config.InternalConfig
}

func NewService(repo IRepository, storage IStorage, grpc_client grpc_auth.AuthServiceClient, cfg *config.InternalConfig) *Service {
	return &Service{
		repo:        repo,
		storage:     storage,
		grpc_client: grpc_client,
		cfg:         cfg,
	}
}

func (s *Service) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	// Password hashing
	var errHash error
	profile.Password, errHash = s.genPasswordHash(profile.Password)
	if errHash != nil {
		return errHash
	}

	//Call database to create profile
	errDb := s.repo.CreateProfile(ctx, profile)
	if errDb != nil {
		return errDb
	}

	// Create new userID
	userId, errUid := uuid.NewV7()
	if errUid != nil {
		return errUid
	}

	// Call Auth by gRPC to sign refresh and access tokens
	result, errCall := s.grpc_client.CreateUser(ctx, &grpc_auth.CreateUserRequest{
		UserId: userId.String(),
	})
	if errCall != nil {
		return errCall
	}

	// Updating entity fields
	profile.UserId = userId
	profile.AccessToken = result.AccessToken
	profile.RefreshToken = result.RefreshToken

	return nil
}
func (s *Service) GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	return s.repo.GetProfile(ctx, userId)
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
		return nil
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
