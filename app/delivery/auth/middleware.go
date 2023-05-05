package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid authorization header")
		}
		tokenString := authorizationHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret123"), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		claims, ok := token.Claims.(*JWTClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
		}
		c.Set("username", claims.Username)
		return next(c)
	}
}
