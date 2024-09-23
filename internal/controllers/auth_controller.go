package controllers

import (
	"net/http"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/services"
	"github.com/ihakalanka/Assessment-Golang-REST-API/pkg"
	"github.com/labstack/echo/v4"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/dtos"
)

func Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	accessToken, refreshToken, err := services.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	claims, err := pkg.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid or expired refresh token")
	}

	accessToken, err := pkg.GenerateAccessToken(claims.UserRole)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not generate access token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": accessToken,
	})
}
