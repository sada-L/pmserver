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

func (cr *cardRepository) Create(card *model.Card) error {
	return nil
}

func (cr *cardRepository) Update(card *model.Card) error {
	return nil
}

func (cr *cardRepository) Delete(id uint) error {
	return nil
}

func (cr *cardRepository) ByUser(ctx context.Context, user *model.User) (*[]model.Card, error) {
	cardQuery := `SELECT id, title, username, password, website, notes, group_id, image, is_favorite FROM cards WHERE user_id = $1`

	var cards []model.Card
	rows, err := cr.q.QueryContext(ctx, cardQuery, user.Id)
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
