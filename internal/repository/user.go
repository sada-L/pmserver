package repository

import (
	"context"
	"time"

	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/postgres"
)

type userRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) model.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    INSERT INTO users (name, email, password)
    VALUES ($1,$2,$3)
    RETURNING id
  `

	return r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.Id)
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	userQuery := `
    SELECT id, name, email, password
    FROM users
    WHERE email = $1
  `

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, userQuery, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}

	return user, nil

	// cardQuery := `
	//    SELECT id, name, username, url, password
	//    FROM cards
	//    WHERE user_id = $1
	//  `
	//
	// rows, err := r.db.QueryContext(ctx, cardQuery, &user.Id)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	//
	// var cards []model.Card
	// for rows.Next() {
	// 	var c model.Card
	// 	if err := rows.Scan(
	// 		&c.Id,
	// 		&c.Name,
	// 		&c.UserName,
	// 		&c.Url,
	// 		&c.Password,
	// 	); err != nil {
	// 		return nil, err
	// 	}
	// 	cards = append(cards, c)
	// }
	//
	// user.Cards = cards
	// return user, nil
}
