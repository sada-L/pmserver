package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/sada-L/pmserver/internal/model"
)

var hmacSampleSecret = []byte("sample-secret")

func GenerateUserToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseUserToken(tokenStr string) (userClaims M, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrUnAuthorized
		}

		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, nil
	}

	return M(claims), nil
}
