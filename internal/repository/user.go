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

func (ur *userRepository) Create(ctx context.Context, user *model.User) (uint, error) {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1,$2,$3) RETURNING id `

	var id uint
	args := []interface{}{user.Username, user.Email, user.PasswordHash}
	err := ur.q.QueryRowContext(ctx, query, args...).Scan(id)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "users_email_key"):
			return 0, model.ErrDuplicateEmail
		case strings.Contains(err.Error(), "users_username_key"):
			return 0, model.ErrDuplicateUsername
		default:
			return 0, err
		}
	}
	return id, nil
}

func (ur *userRepository) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3`

	args := []interface{}{user.Username, user.Email, user.Id}
	if err := ur.q.QueryRowContext(ctx, query, args...).Err(); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = $1`

	if err := ur.q.QueryRowContext(ctx, query, id).Err(); err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) ById(ctx context.Context, id uint) (*model.User, error) {
	userQuery := `SELECT id, username, email, password_hash FROM users WHERE id = $1`

	user := &model.User{}
	err := ur.q.QueryRowContext(ctx, userQuery, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) ByEmail(ctx context.Context, email string) (*model.User, error) {
	userQuery := ` SELECT id, username, email, password_hash FROM users WHERE email = $1 `

	user := &model.User{}
	err := ur.q.QueryRowContext(ctx, userQuery, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}
