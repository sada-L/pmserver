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

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", fmt.Errorf("token - GenUserToken - SignedString: %w", err)
	}

	return tokenString, nil
}

func ParseUserToken(tokenStr string) (userClaims M, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrUnAuthorized
		}

		return []byte(os.Getenv("SECRET")), nil
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
