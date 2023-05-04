package auth

import (
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, h *AuthHandler) {
	authGroup := e.Group("/auth")

	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
}
