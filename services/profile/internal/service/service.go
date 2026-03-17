package service

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
)

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
	repo    IRepository
	storage IStorage
}

func NewService(repo IRepository, storage IStorage) *Service {
	return &Service{repo: repo, storage: storage}
}

func (s *Service) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	errDb := s.repo.CreateProfile(ctx, profile)
	if errDb != nil {
		return errDb
	}

	// s.storage.PutObject()
	return nil
}
func (s *Service) GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error) {
	return s.repo.GetProfile(ctx, userId)
}
func (s *Service) UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error) {
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
