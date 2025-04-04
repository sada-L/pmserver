package service

import (
	"context"
	"time"

	"github.com/sada-L/pmserver/internal/model"
)

type userService struct {
	UserRepository model.UserRepository
}

func NewUserService(r model.UserRepository) model.UserService {
	return &userService{UserRepository: r}
}

func (s *userService) Create(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.UserRepository.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, user *model.User) error {
	return s.UserRepository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.UserRepository.Delete(ctx, id)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.UserRepository.GetByEmail(ctx, email)
}
