package api

import (
	"mime/multipart"

	"github.com/google/uuid"
)

// Request
type RequestCreateaProfile struct {
	Username                 string                 `form:"username"`
	Name                     string                 `form:"name"`
	Surname                  string                 `form:"surname"`
	Email                    string                 `form:"email"`
	Password                 string                 `form:"password"`
	Description              string                 `form:"description"`
	Latitude                 float64                `form:"latitude"`
	Longitude                float64                `form:"longitude"`
	PreferencesMaxDistMeters int                    `form:"max_dist_meters"`
	FavArtStyles             []string               `form:"fav_art_styles"`
	Photos                   []multipart.FileHeader `form:"photos"`
}

type RequestGetProfile struct {
	UserId uuid.UUID `json:"user_id"`
}

type RequestUpdateProfile struct {
	Username    *string  `json:"username"`
	Name        *string  `json:"name"`
	Surname     *string  `json:"surname"`
	Email       *string  `json:"email"`
	Password    *string  `json:"password"`
	Description *string  `json:"description"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
}

// Response
type ResponseProfile struct {
	Username    string  `json:"username"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type ResponseCreateProfile struct {
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	Description  string  `json:"description"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}
