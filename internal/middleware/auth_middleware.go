package middleware

import (
	"net/http"
	"strings"
	"github.com/labstack/echo/v4"
	"github.com/ihakalanka/Assessment-Golang-REST-API/pkg"
)

func RoleMiddleware(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, "missing or invalid token")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := pkg.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			}

			if claims.UserRole != requiredRole {
				return c.JSON(http.StatusForbidden, "insufficient permissions")
			}

			return next(c)
		}
	}
}
