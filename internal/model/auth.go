package model

import (
	"github.com/golang-jwt/jwt"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}
