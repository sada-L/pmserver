package service

import (
	"context"
	"fmt"

	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/internal/repository"
	"github.com/sada-L/pmserver/pkg/postgres"
)

type userService struct {
	db *postgres.DB
}

func NewUserService(db *postgres.DB) model.UserService {
	return &userService{db: db}
}

func (us *userService) Authenticate(ctx context.Context, email, password string) (*model.User, error) {
	uc := repository.NewUserRepository(us.db)
	user, err := uc.UserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("userService - Authenticate - ur.UserByEmail: %w", err)
	}

	if !user.VerifyPassword(password) {
		return nil, model.ErrUnAuthorized
	}

	return user, nil
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) error {
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("userService - CreateUser - us.db.BeginTx: %w", err)
	}

	defer tx.Rollback()

	ur := repository.NewUserRepository(tx)
	if err := ur.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("userService - CreateUser - ur.CreateUser: %w", err)
	}

	return tx.Commit()
}

func (us *userService) UpdateUser(ctx context.Context, user *model.User) error {
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("userService - UpdateUser - us.db.BeginTx: %w", err)
	}

	defer tx.Rollback()

	ur := repository.NewUserRepository(tx)
	if err := ur.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("userService - UpdateUser - ur.UpdateUser: %w", err)
	}

	return tx.Commit()
}

func (us *userService) DeleteUser(ctx context.Context, id uint) error {
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("userService - DeleteUser - us.db.BeginTx: %w", err)
	}

	defer tx.Rollback()

	ur := repository.NewUserRepository(tx)
	if err := ur.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("userService - DeleteUser - ur.DeleteUser: %w", err)
	}

	return tx.Commit()
}

func (us *userService) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	ur := repository.NewUserRepository(us.db)
	user, err := ur.UserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("userService - UserByEmail - ur.UserByEmail: %w", err)
	}

	return user, nil
}
