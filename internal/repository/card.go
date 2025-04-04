package repository

import (
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
