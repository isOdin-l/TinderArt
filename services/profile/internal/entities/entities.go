package entities

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type Profile struct {
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

	// Photos
	PhotoUrls  []string
	PhotosIds  []uuid.UUID
	PhotoFiles []multipart.FileHeader

	//Preferences
	PreferencesMaxDistMeters int

	// Favourite art style
	FavArtStylesIds []uuid.UUID
	FavArtStyles    []string
}

type UpdateProfile struct {
	UserId uuid.UUID

	// Optional fields
	Username    *string
	Name        *string
	Surname     *string
	Email       *string
	Password    *string
	Description *string
	Latitude    *float64
	Longitude   *float64
}
