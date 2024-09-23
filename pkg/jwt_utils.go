package pkg

import (
	"time"
	"github.com/golang-jwt/jwt"
	"os"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/dtos"
	"errors"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateAccessToken(userRole string) (string, error) {
	claims := &dto.JWTClaims{
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Assessment-Golang-REST-API",
			Subject:   "Access Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userRole string) (string, error) {
	claims := &dto.JWTClaims{
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Assessment-Golang-REST-API",
			Subject:   "Refresh Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*dto.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*dto.JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func ValidateRefreshToken(refreshToken string) (*dto.JWTClaims, error) {
	token, _ := jwt.ParseWithClaims(refreshToken, &dto.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*dto.JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid or expired refresh token")
}