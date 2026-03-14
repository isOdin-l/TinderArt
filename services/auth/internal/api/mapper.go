package entities

import (
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"github.com/isOdin-l/TinderArt/services/auth/pkg/api"
)

// Function design:
// FromAPI<RequestModel>To<EntityName>  (req *grpc_gen.<RequestModel>) *<EntityName>
// From<EntityName>ToAPI<ResponseModel> (req *<EntityName>) *grpc_gen.<ResponseModel>

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
