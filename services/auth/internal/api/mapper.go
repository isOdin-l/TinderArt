package entities

import (
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"github.com/isOdin-l/TinderArt/services/auth/pkg/api"
)

// Function design:
// FromAPI<RequestModel>To<EntityName>  (req *api.<RequestModel>) *<EntityName>
// From<EntityName>ToAPI<ResponseModel> (req *<EntityName>) *api.<ResponseModel>

// --------- Request ---------
func FromAPIValidateTokenToValidateToken(req *api.ValidateToken) *entities.ValidateToken {
	return &entities.ValidateToken{
		AccessToken: req.AccessToken,
	}
}

func FromAPILoginToLogin(req *api.Login) *entities.Login {
	return &entities.Login{
		Username: req.Username,
		Password: req.Password,
	}
}

func FromAPIRegistrationToRegistration(req *api.Registration) *entities.Registration {
	return &entities.Registration{
		Username:    req.Username,
		Name:        req.Name,
		Surname:     req.Surname,
		Email:       req.Email,
		Password:    req.Password,
		Description: req.Description,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
	}
}

func FromAPIRefreshTokenToRefreshToken(req *api.RefreshAccessToken) *entities.RefreshAccessToken {
	return &entities.RefreshAccessToken{
		RefreshToken: req.RefreshToken,
	}
}

// --------- Response ---------
func FromLoginToTokenResponse(login *entities.Login) *api.TokenResponse {
	return &api.TokenResponse{
		AccessToken:  login.AccessToken,
		RefreshToken: login.AccessToken,
	}
}

func FromAuthResultToTokenResponse(result *entities.AuthResult) *api.TokenResponse {
	return &api.TokenResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}
}

func FromRegistrationToTokenResponse(reg *entities.Registration) *api.TokenResponse {
	return &api.TokenResponse{
		AccessToken:  reg.AccessToken,
		RefreshToken: reg.RefreshToken,
	}
}

func FromRefreshAccessTokenToTokenResponse(refresh *entities.RefreshAccessToken) *api.TokenResponse {
	return &api.TokenResponse{
		AccessToken:  refresh.AccessToken,
		RefreshToken: refresh.RefreshToken,
	}
}
