package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	mapper "github.com/isOdin-l/TinderArt/services/profile/internal/api"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/isOdin-l/TinderArt/services/profile/pkg/api"
	"github.com/labstack/echo/v5"
)

type IService interface {
	CreateProfile(ctx context.Context, profile *entities.Profile) error
	GetProfile(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)
	UpdateProfile(ctx context.Context, profile *entities.UpdateProfile) (*entities.Profile, error)
	DeleteProfile(ctx context.Context, userId uuid.UUID) error
	GetStack(ctx context.Context, userId uuid.UUID) (*entities.Profile, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateProfile(c *echo.Context) error {
	var req api.RequestCreateaProfile
	if errBind := c.Bind(&req); errBind != nil {
		slog.Error(fmt.Sprintf("error: %s data: request data", errBind.Error()))
		return c.JSON(http.StatusBadRequest, "invalid data")
	}

	entity := mapper.FromAPICreateProfileToEntity(&req)

	errService := h.service.CreateProfile(c.Request().Context(), entity)
	if errService != nil {
		slog.Error(fmt.Sprintf("error: %s request data", errService.Error()))
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, mapper.FromEntityToAPICreateProfile(entity))
}

func (h *Handler) GetProfile(c *echo.Context) error {
	var req api.RequestGetProfile
	if errBind := c.Bind(&req); errBind != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errBind.Error(), req))
		return c.JSON(http.StatusBadRequest, "invalid data")
	}

	// Mapping string to uuid
	entity, errMap := mapper.FromAPIGetProfileToEntity(&req)
	if errMap != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errMap.Error(), entity))
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	profile, errService := h.service.GetProfile(c.Request().Context(), entity)
	if errService != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errService.Error(), profile))
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, mapper.FromEntityToAPIGetProfile(profile))
}

func (h *Handler) UpdateProfile(c *echo.Context) error {
	var req api.RequestUpdateProfile
	if errBind := c.Bind(&req); errBind != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errBind.Error(), req))
		return c.JSON(http.StatusBadRequest, "invalid data")
	}

	// Move from api model to entity and
	// get userId from context
	entity := mapper.FromAPIUpdateProfileToEntity(&req)
	entity.UserId = c.Get("user_id").(uuid.UUID)

	profile, errService := h.service.UpdateProfile(c.Request().Context(), entity)
	if errService != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errService.Error(), profile))
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, mapper.FromEntityToAPIGetProfile(profile)) // return profile data
}

func (h *Handler) DeleteProfile(c *echo.Context) error {
	// Get from context userId
	userId := c.Get("user_id").(uuid.UUID)

	if errServ := h.service.DeleteProfile(c.Request().Context(), userId); errServ != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errServ.Error(), userId))
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetStack(c *echo.Context) error {
	userId := c.Get("user_id").(uuid.UUID)

	entity, errServer := h.service.GetStack(c.Request().Context(), userId)
	if errServer != nil {
		slog.Error(fmt.Sprintf("error: %s data:%s", errServer.Error(), entity))
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, mapper.FromEntityToAPIGetStack(entity))
}
