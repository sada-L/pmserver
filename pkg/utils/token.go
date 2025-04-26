package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sada-L/pmserver/internal/model"
)

func CreateAccessToken(user *model.User, expiry int) (string, error) {
	exp := time.Now().Add(time.Duration(expiry) * time.Hour).Unix()
	claims := &model.AccessTokenClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getSecret())
	if err != nil {
		return "", fmt.Errorf("token - CreateUserToken - SignedString: %w", err)
	}

	return tokenString, nil
}

func CreateRefreshToken(user *model.User, expiry int) (string, error) {
	exp := time.Now().Add(time.Duration(expiry) * time.Hour * 24).Unix()
	claims := &model.RefreshTokenClaims{
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getSecret())
	if err != nil {
		return "", fmt.Errorf("token - CreateUserToken - SignedString: %w", err)
	}

	return tokenString, nil
}

func ParseUserToken(requestToken string) (M, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrUnAuthorized
		}
		return getSecret(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token - ParseUserToken - jwt.Parse: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return M(claims), nil
}

func NewAuthResponse(user *model.User, expire int) (*model.AuthResponse, error) {
	accessToken, err := CreateAccessToken(user, expire)
	if err != nil {
		return nil, err
	}

	refreshToken, err := CreateRefreshToken(user, expire)
	if err != nil {
		return nil, err
	}

	authResponse := &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return authResponse, nil
}

func getSecret() []byte {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		secret = []byte("JWT_SECRET")
	}
	return secret
}
