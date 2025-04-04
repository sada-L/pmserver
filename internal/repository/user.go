package repository

import (
	"context"
	"strings"

	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/postgres"
)

type userRepository struct {
	q postgres.Querier
}

func NewUserRepository(q postgres.Querier) model.UserRepository {
	return &userRepository{q: q}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `
    INSERT INTO users (username, email, password_hash)
    VALUES ($1,$2,$3) RETURNING id
  `
	args := []interface{}{user.Username, user.Email, user.PasswordHash}
	err := ur.q.QueryRowContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "users_email_key"):
			return model.ErrDuplicateEmail
		case strings.Contains(err.Error(), "users_username_key"):
			return model.ErrDuplicateUsername
		default:
			return err
		}
	}
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	return nil
}

func (r *userRepository) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	userQuery := `
    SELECT id, username, email, password_hash
    FROM users
    WHERE email = $1
  `

	user := &model.User{}
	err := r.q.QueryRowContext(ctx, userQuery, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
