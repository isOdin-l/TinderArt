package api

// REQUEST
type Registration struct {
	Username    string  `json:"username"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshAccessToken struct {
	RefreshToken string `json:"refresh_token"`
}

type ValidateToken struct {
	AccessToken string `header:"Authorization"`
}

// RESPONSE
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ValidateResponse struct {
	Valid bool `json:"valid"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
