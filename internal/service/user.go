package service

import (
	"context"

	"github.com/sada-L/pmserver/internal/model"
)

type userService struct {
	ur model.UserRepository
}

func NewUserService(ur model.UserRepository) model.UserService {
	return &userService{ur: ur}
}

func (us *userService) Authenticate(ctx context.Context, email, password string) (*model.User, error) {
	return nil, nil
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) error {
	return us.ur.CreateUser(ctx, user)
}

func (us *userService) UpdateUser(ctx context.Context, user *model.User) error {
	return us.ur.UpdateUser(ctx, user)
}

func (us *userService) DeleteUser(ctx context.Context, id uint) error {
	return us.ur.DeleteUser(ctx, id)
}

func (us *userService) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	return us.ur.UserByEmail(ctx, email)
}
