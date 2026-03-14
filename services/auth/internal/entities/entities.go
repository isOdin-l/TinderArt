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
	Username    string
	Name        string
	Surname     string
	Email       string
	Password    string
	Description string
	Latitude    float64
	Longitude   float64
}

type Login struct {
	UserId          uuid.UUID
	Username        string
	Password        string
	NewRefreshToken string
	NewAccessToken  string
}

type RefreshAccessToken struct {
	RefreshToken string
}
