package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	config "github.com/isOdin-l/TinderArt/services/auth/configs"
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type IRepo interface {
	GetUser() (*entities.Login, error) //
	GetRefreshToken()
}

type AuthService struct {
	cfg  *config.InternalConfig
	repo IRepo
}

func NewService(cfg *config.InternalConfig, repo IRepo) *AuthService {
	return &AuthService{cfg: cfg, repo: repo}
}

func (s *AuthService) Registrations(ctx context.Context, entity *entities.Registration) error {
	return nil
}

func (s *AuthService) Login(ctx context.Context, entity *entities.Login) error {
	userDb, errDb := s.repo.GetUser() // send username and hashed password
	if errDb != nil {
		return nil
	}

	s.signAccessToken(userDb.UserId)

	return nil
}
func (s *AuthService) RefreshAccessToken(ctx context.Context, entity *entities.RefreshAccessToken) error {
	return nil
}
func (s *AuthService) ValidateAccessToken(ctx context.Context, entity *entities.ValidateToken) error {
	valRes, errVal := s.validate(entity.AccessToken)
	if errVal != nil {
		return errVal
	}

	if valRes {
		return nil
	}

	return errors.New("Unauthorized") // Create custom errors
}

func (s *AuthService) signAccessToken(userId uuid.UUID) (string, error) {
	claims := &entities.JwtTokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.TokenTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString([]byte(s.cfg.JwtSignKey))
}

func (s *AuthService) signRefreshToken(userId uuid.UUID) (string, error) {
	claims := &entities.JwtTokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.TokenTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString([]byte(s.cfg.JwtSignKey))
}

func (s *AuthService) validate(accessToken string) (bool, error) {
	token, errParse := jwt.ParseWithClaims(accessToken, &entities.JwtTokenClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodES256 {
				return nil, jwt.ErrInvalidKeyType
			}

			return []byte(s.cfg.JwtSignKey), nil
		})

	return token.Valid, errParse
}

func (s *AuthService) genPasswordHash(password string) (string, error) {
	hash, errHash := bcrypt.GenerateFromPassword([]byte(password), s.cfg.HashMinCost)
	return string(hash), errHash
}

func (s *AuthService) grpcCallCreateUser() {} // gRPCS work to connect with Profile-service
