package model

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint   `json:"id"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"-"`
}

type UserRepository interface {
	Create(context.Context, *User) (uint, error)
	Update(context.Context, *User) error
	Delete(context.Context, uint) error
	ById(context.Context, uint) (*User, error)
	ByEmail(context.Context, string) (*User, error)
}

type UserService interface {
	CreateUser(context.Context, *User) (uint, error)
	UpdateUser(context.Context, *User) error
	DeleteUser(context.Context, uint) error
	UserById(context.Context, uint) (*User, error)
	UserByEmail(context.Context, string) (*User, error)
}

func (u *User) SetPassword(password string) error {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("user - SetPassword - GenHash: %w", err)
	}
	u.PasswordHash = string(hashBytes)

	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}
