package entities

import (
	"github.com/google/uuid"
	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	"github.com/isOdin-l/TinderArt/services/auth/internal/entities"
	"github.com/isOdin-l/TinderArt/services/auth/pkg/api"
)

// Function design:
// FromAPI<RequestModel>To<EntityName>  (req *api.<RequestModel>) *<EntityName>
// From<EntityName>ToAPI<ResponseModel> (req *<EntityName>) *api.<ResponseModel>

// --------- Request ---------
func FromAPIValidateTokenToValidateToken(req *grpc_auth.ValidateRequest) *entities.ValidateToken {
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

func FromAPIRegistrationToRegistration(req *grpc_auth.CreateUserRequest) *entities.Registration {
	return &entities.Registration{
		UserId: uuid.MustParse(req.GetUserId()),
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

func FromRegistrationToTokenResponse(req *entities.Registration) *grpc_auth.CreateUserResponse {
	return &grpc_auth.CreateUserResponse{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	}
}

func FromRefreshAccessTokenToTokenResponse(refresh *entities.RefreshAccessToken) *api.TokenResponse {
	return &api.TokenResponse{
		AccessToken:  refresh.AccessToken,
		RefreshToken: refresh.RefreshToken,
	}
}
