package repository

import (
	"context"
	"database/sql"

	"github.com/sada-L/pmserver/internal/model"
)

type cardRepository struct {
	db *sql.DB
}

func NewCardRepository(db *sql.DB) model.CardRepository {
	return &cardRepository{db: db}
}

func (r cardRepository) Create(card *model.Card) error {
	return nil
}

func (r cardRepository) Update(card *model.Card) error {
	return nil
}

func (r cardRepository) Delete(id uint) error {
	return nil
}

func (r *cardRepository) ByUserId(ctx context.Context, userId string) (*[]model.Card, error) {
	cardQuery := `SELECT * FROM cards WHERE user_id = $1`

	var cards []model.Card
	rows, err := r.db.QueryContext(ctx, cardQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var card model.Card
		err := rows.Scan(card.Id, card.Name, card.UserName, card.GroupId, card.Password, card.Url)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &cards, nil
}
