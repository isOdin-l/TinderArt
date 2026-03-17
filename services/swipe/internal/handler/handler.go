package handler

import (
	"context"
	"net/http"

	mapper "github.com/isOdin-l/TinderArt/services/swipe/internal/api"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
	"github.com/isOdin-l/TinderArt/services/swipe/pkg/api"
	"github.com/labstack/echo/v5"
)

type IService interface {
	CreateSwipe(ctx context.Context, swipe *entities.Swipe) error
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateSwipe(c *echo.Context) error {
	var req api.CreateSwipeRequest
	if errBind := c.Bind(&req); errBind != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	entity := mapper.FromApiSwipeToEntity(&req)
	errService := h.service.CreateSwipe(c.Request().Context(), entity)
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.NoContent(http.StatusOK)
}
