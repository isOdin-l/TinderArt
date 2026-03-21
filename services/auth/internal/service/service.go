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

var (
	accessSignType  = jwt.SigningMethodHS256
	refreshSignType = jwt.SigningMethodHS512
)

type IRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*entities.Login, error)
	SaveRefreshToken(ctx context.Context, userId uuid.UUID, refreshToken string) error
}

type TxManager interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error)
}

type AuthService struct {
	cfg  *config.InternalConfig
	repo IRepo
	txm  TxManager
}

func NewService(cfg *config.InternalConfig, repo IRepo, txm TxManager) *AuthService {
	return &AuthService{cfg: cfg, repo: repo, txm: txm}
}

func (s *AuthService) Registrations(ctx context.Context, entity *entities.Registration) error {
	var errSign error

	// Generate tokens
	entity.AccessToken, errSign = s.signAccessToken(entity.UserId)
	if errSign != nil {
		return errSign
	}
	entity.RefreshToken, errSign = s.signRefreshToken(entity.UserId)
	if errSign != nil {
		return errSign
	}

	// Save refresh token
	return s.repo.SaveRefreshToken(ctx, entity.UserId, entity.RefreshToken)
}

func (s *AuthService) Login(ctx context.Context, entity *entities.Login) error {
	_, errTx := s.txm.WithTx(ctx, func(ctx context.Context) (any, error) {
		// Check if user exist
		userDb, errRepo := s.repo.GetUserByUsername(ctx, entity.Username)
		if errRepo != nil {
			return nil, errRepo
		}

		// Verify password
		if errCompare := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(entity.Password)); errCompare != nil {
			return nil, errCompare
		}

		// Generate tokens
		var errTokens error
		entity.AccessToken, errTokens = s.signAccessToken(userDb.UserId)
		if errTokens != nil {
			return nil, errTokens
		}

		entity.RefreshToken, errTokens = s.signRefreshToken(userDb.UserId)
		if errTokens != nil {
			return nil, errTokens
		}

		entity.UserId = userDb.UserId
		// Save refresh token
		return nil, s.repo.SaveRefreshToken(ctx, entity.UserId, entity.RefreshToken)
	})
	return errTx
}
func (s *AuthService) RefreshAccessToken(ctx context.Context, entity *entities.RefreshAccessToken) error {
	// Parse refresh token to get claims
	token, err := s.parseRefreshToken(entity.RefreshToken)
	if err != nil {
		return errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*entities.JwtTokenClaims)
	if !ok || !token.Valid {
		return errors.New("invalid refresh token")
	}

	// Generate new access token
	entity.AccessToken, err = s.signAccessToken(claims.UserId)

	return err
}
func (s *AuthService) ValidateAccessToken(ctx context.Context, entity *entities.ValidateToken) (uuid.UUID, error) {
	token, errVal := s.parseAccessToken(entity.AccessToken)
	if errVal != nil {
		return uuid.Nil, errVal
	}

	claims, ok := token.Claims.(*entities.JwtTokenClaims)
	if !ok || !token.Valid {
		return uuid.Nil, errors.New("Unauthorized")
	}

	return claims.UserId, nil
}

// Internal methods
func (s *AuthService) signRefreshToken(userId uuid.UUID) (string, error) {
	claims := &entities.JwtTokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.RefreshTokenTTL)),
		},
	}
	token := jwt.NewWithClaims(refreshSignType, claims)

	return token.SignedString([]byte(s.cfg.RefreshSignKey))
}

func (s *AuthService) parseRefreshToken(refreshToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(refreshToken, &entities.JwtTokenClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method != refreshSignType {
				return nil, jwt.ErrInvalidKeyType
			}

			return []byte(s.cfg.RefreshSignKey), nil
		})
}

func (s *AuthService) signAccessToken(userId uuid.UUID) (string, error) {
	claims := &entities.JwtTokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.AccessTokenTTL)),
		},
	}
	token := jwt.NewWithClaims(accessSignType, claims)

	return token.SignedString([]byte(s.cfg.AccessSignKey))
}

func (s *AuthService) parseAccessToken(accessToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(accessToken, &entities.JwtTokenClaims{},
		func(t *jwt.Token) (any, error) {
			if t.Method != accessSignType {
				return nil, jwt.ErrInvalidKeyType
			}

			return []byte(s.cfg.AccessSignKey), nil
		})
}
