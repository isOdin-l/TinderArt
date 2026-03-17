package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	mapper "github.com/isOdin-l/TinderArt/services/profile/internal/api"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/isOdin-l/TinderArt/services/profile/pkg/api"
	"github.com/labstack/echo/v5"
)

type IService interface {
	GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)
	UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error)
	DeleteProfile(ctx context.Context, userId uuid.UUID) error
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetProfile(c *echo.Context) error {
	var req api.RequestGetProfile
	if errBind := c.Bind(&req); errBind != nil {
		return c.JSON(http.StatusBadRequest, "invalid data")
	}

	profile, errService := h.service.GetProfile(c.Request().Context(), mapper.FromAPIGetProfileToEntity(&req))
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, mapper.FromEntityToAPIGetProfile(profile))
}

func (h *Handler) UpdateProfile(c *echo.Context) error {
	var req api.RequestUpdateProfile
	if errBind := c.Bind(&req); errBind != nil {
		return c.JSON(http.StatusBadRequest, "invalid data")
	}
	entity := mapper.FromAPIUpdateProfileToEntity(&req)
	entity.UserId = c.Get("user_id").(uuid.UUID)

	profile, errService := h.service.UpdateProfile(c.Request().Context(), entity)
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, mapper.FromEntityToAPIGetProfile(profile)) // return profile data
}

func (h *Handler) DeleteProfile(c *echo.Context) error {
	// Get from context userId
	userId := c.Get("user_id").(uuid.UUID)

	h.service.DeleteProfile(c.Request().Context(), userId)

	return c.NoContent(http.StatusOK)
}
