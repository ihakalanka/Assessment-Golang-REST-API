package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/ihakalanka/Assessment-Golang-REST-API/utils"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "missing token")
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.JSON(http.StatusUnauthorized, "invalid token format")
		}

		tokenString := authHeader[7:]

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "invalid or expired token")
		}

		c.Set("userRole", claims.UserRole)
		return next(c)
	}
}