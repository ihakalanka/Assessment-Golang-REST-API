package services

import (
	"errors"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/config"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/models"
	"github.com/ihakalanka/Assessment-Golang-REST-API/pkg"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(email, password string) (string, string, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", "", errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := pkg.GenerateAccessToken(user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateRefreshToken(user.Role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}