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

func (us *userService) CreateUser(ctx context.Context, user *model.User) (uint, error) {
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("userService - CreateUser - us.db.BeginTx: %w", err)
	}

	defer tx.Rollback()

	ur := repository.NewUserRepository(tx)
	id, err := ur.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("userService - CreateUser - ur.CreateUser: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("userService - CreateUser - tx.Commit: %w", err)
	}

	return id, nil
}

func (us *userService) UpdateUser(ctx context.Context, user *model.User) error {
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("userService - UpdateUser - us.db.BeginTx: %w", err)
	}

	defer tx.Rollback()

	ur := repository.NewUserRepository(tx)
	if err = ur.Update(ctx, user); err != nil {
		return fmt.Errorf("userService - UpdateUser - ur.UpdateUser: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("userService - UpdateUser - tx.Commit: %w", err)
	}

	return nil
}

func (us *userService) DeleteUser(ctx context.Context, id uint) error {
	tx, err := us.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("userService - DeleteUser - us.db.BeginTx: %w", err)
	}

	defer tx.Rollback()

	ur := repository.NewUserRepository(tx)
	if err = ur.Delete(ctx, id); err != nil {
		return fmt.Errorf("userService - DeleteUser - ur.DeleteUser: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("userService - DeleteUser - tx.Commit: %w", err)
	}

	return nil
}

func (us *userService) UserById(ctx context.Context, id uint) (*model.User, error) {
	ur := repository.NewUserRepository(us.db)
	user, err := ur.ById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService - UserById - ur.ById: %w", err)
	}

	return user, nil
}

func (us *userService) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	ur := repository.NewUserRepository(us.db)
	user, err := ur.ByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("userService - UserByEmail - ur.UserByEmail: %w", err)
	}

	return user, nil
}
