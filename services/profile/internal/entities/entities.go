package entities

import "github.com/google/uuid"

type Profile struct {
	UserId      uuid.UUID
	Username    string
	Name        string
	Surname     string
	Email       string
	Password    string
	Description string
	Latitude    float64
	Longitude   float64
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
