package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/sada-L/pmserver/internal/model"
)

func GenerateUserToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})

	key := []byte(os.Getenv("JWT_KEY"))
	if len(key) == 0 {
		key = []byte("JWT_KEY")
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("token - GenUserToken - SignedString: %w", err)
	}

	return tokenString, nil
}

func ParseUserToken(tokenStr string) (userClaims M, err error) {
	key := []byte(os.Getenv("JWT_KEY"))
	if len(key) == 0 {
		key = []byte("JWT_KEY")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrUnAuthorized
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token - ParsUserToken - jwt.Parse: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil
	}

	return M(claims), nil
}
