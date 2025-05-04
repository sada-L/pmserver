package repository

import (
	"context"
	"github.com/sada-L/pmserver/pkg/postgres"

	"github.com/sada-L/pmserver/internal/model"
)

type cardRepository struct {
	q postgres.Querier
}

func NewCardRepository(q postgres.Querier) model.CardRepository {
	return &cardRepository{q: q}
}

func (cr *cardRepository) Create(ctx context.Context, card *model.Card) (uint, error) {
	query := `INSERT INTO cards (title, username, password, website, notes, group_id, image, is_favorite, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	args := []interface{}{card.Title, card.Username, card.Password, card.Website, card.Notes, card.GroupId, card.Image, card.IsFavorite, card.UserId}
	if err := cr.q.QueryRowContext(ctx, query, args...).Scan(&card.Id); err != nil {
		return 0, err
	}

	return card.Id, nil
}

func (cr *cardRepository) Update(ctx context.Context, card *model.Card) error {
	query := ` UPDATE cards SET title = $1, username = $2, password = $3, website = $4, notes = $5, group_id = $6, image = $7, is_favorite = $8 WHERE id = $9`

	args := []interface{}{card.Title, card.Username, card.Password, card.Website, card.Notes, card.GroupId, card.Image, card.IsFavorite, card.Id}
	if err := cr.q.QueryRowContext(ctx, query, args...).Err(); err != nil {
		return err
	}

	return nil
}

func (cr *cardRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM cards WHERE id = $1`

	if err := cr.q.QueryRowContext(ctx, query, id).Err(); err != nil {
		return err
	}

	return nil
}

func (cr *cardRepository) ByUser(ctx context.Context, user *model.User) (*[]model.Card, error) {
	query := `SELECT id, title, username, password, website, notes, group_id, image, is_favorite FROM cards WHERE user_id = $1`

	var cards []model.Card
	rows, err := cr.q.QueryContext(ctx, query, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var card model.Card
		if err = rows.Scan(
			&card.Id,
			&card.Title,
			&card.Username,
			&card.Password,
			&card.Website,
			&card.Notes,
			&card.GroupId,
			&card.Image,
			&card.IsFavorite); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &cards, nil
}
