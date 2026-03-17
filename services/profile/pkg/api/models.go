package api

import "github.com/google/uuid"

// Request
type RequestCreateaProfile struct {
	Username    string  `json:"username"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
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
