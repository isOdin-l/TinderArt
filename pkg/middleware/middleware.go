package middleware

import (
	"net/http"
	"strings"

	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	"github.com/labstack/echo/v5"
)

const UserIdContextKey string = "user_id"

type Middleware struct {
	auth grpc_auth.AuthServiceClient
}

func NewMiddleware(client grpc_auth.AuthServiceClient) *Middleware {
	return &Middleware{
		auth: client,
	}
}
func (m *Middleware) Validation() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "missing authorization header",
				})
			}

			// delete Bearer part token
			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid token",
				})
			}

			// gRPC call
			response, errVal := m.auth.Validate(c.Request().Context(), &grpc_auth.ValidateRequest{
				AccessToken: token,
			})

			if errVal != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid token",
				})
			}

			// put uesr_id to context
			c.Set(UserIdContextKey, response.UserId)

			return next(c)
		}
	}
}
