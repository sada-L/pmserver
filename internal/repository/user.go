package repository

import (
	"context"

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
    INSERT INTO users (name, email, password)
    VALUES ($1,$2,$3) RETURNING id
  `

	args := []interface{}{user.Name, user.Email, user.Password}
	return ur.q.QueryRowContext(ctx, query, args...).Scan(&user.Id)
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	return nil
}

func (r *userRepository) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	userQuery := `
    SELECT id, name, email, password
    FROM users
    WHERE email = $1
  `

	user := &model.User{}
	err := r.q.QueryRowContext(ctx, userQuery, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
