package model

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint     `json:"-"`
	Username     string   `json:"username,omitempty"`
	Email        string   `json:"email,omitempty"`
	Token        string   `json:"token,omitempty"`
	PasswordHash string   `json:"-"`
	Cards        []*Card  `json:"-"`
	Groups       []*Group `json:"-"`
}

type UserRepository interface {
	CreateUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
	DeleteUser(context.Context, uint) error
	UserByEmail(context.Context, string) (*User, error)
}

type UserService interface {
	Authenticate(ctx context.Context, email, password string) (*User, error)
	CreateUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
	DeleteUser(context.Context, uint) error
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
