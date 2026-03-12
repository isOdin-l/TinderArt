package service

import config "github.com/isOdin-l/TinderArt/services/auth/configs"

type IRepo interface {
}

type AuthService struct {
	cfg  *config.InternalConfig
	repo IRepo
}

func NewService(cfg *config.InternalConfig, repo IRepo) *AuthService {
	return &AuthService{cfg: cfg, repo: repo}
}
