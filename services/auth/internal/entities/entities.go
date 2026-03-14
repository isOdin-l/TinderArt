package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtTokenClaims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

type ValidateToken struct {
	AccessToken string
}

type Registration struct {
	UserId       uuid.UUID
	Username     string
	Name         string
	Surname      string
	Email        string
	Password     string
	Description  string
	Latitude     float64
	Longitude    float64
	AccessToken  string
	RefreshToken string
}

type Login struct {
	UserId       uuid.UUID
	Username     string
	Password     string
	RefreshToken string
	AccessToken  string
}

type RefreshAccessToken struct {
	RefreshToken string
	AccessToken  string
}

type AuthResult struct {
	UserId       uuid.UUID
	AccessToken  string
	RefreshToken string
}
