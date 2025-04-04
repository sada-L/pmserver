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
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uint) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserService interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uint) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
