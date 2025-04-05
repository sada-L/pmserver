package model

import (
	"context"
)

type User struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Cards    []Card  `json:"cards"`
	Groups   []Group `json:"groups"`
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
