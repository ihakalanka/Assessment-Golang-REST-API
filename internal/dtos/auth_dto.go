package dto

import (
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	UserRole string `json:"role"`
	jwt.StandardClaims
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}