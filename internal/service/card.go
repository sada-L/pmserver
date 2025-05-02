package service

import (
	"context"
	"fmt"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/internal/repository"
	"github.com/sada-L/pmserver/pkg/postgres"
)

type cardService struct {
	db *postgres.DB
}

func NewCardService(db *postgres.DB) model.CardService {
	return &cardService{db: db}
}

func (cs *cardService) CreateCard(ctx context.Context, card *model.Card) (uint, error) {
	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("cardService - CreateCard - cs.db.BeginTx: %w", err)
	}
	defer tx.Rollback()

	cr := repository.NewCardRepository(tx)
	id, err := cr.Create(ctx, card)
	if err != nil {
		return 0, fmt.Errorf("cardService - CreateCard - cr.Create: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("cardService - CreateCard - tx.Commit: %w", err)
	}

	return id, err
}

func (cs *cardService) UpdateCard(ctx context.Context, card *model.Card) error {
	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("cardService - UpdateCard - cs.db.BeginTx: %w", err)
	}
	defer tx.Rollback()

	cr := repository.NewCardRepository(tx)
	if err := cr.Update(ctx, card); err != nil {
		return fmt.Errorf("cardService - UpdateCard - cr.Update: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cardService - UpdateCard - tx.Commit: %w", err)
	}

	return nil
}

func (cs *cardService) DeleteCard(ctx context.Context, id uint) error {
	tx, err := cs.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("cardService - DeleteCard - cs.db.BeginTx: %w", err)
	}
	defer tx.Rollback()

	cr := repository.NewCardRepository(tx)
	if err := cr.Delete(ctx, id); err != nil {
		return fmt.Errorf("cardService - DeleteCard - cr.Delete: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("cardService - DeleteCard - tx.Commit: %w", err)
	}

	return nil
}

func (cs *cardService) CardsByUser(ctx context.Context, user *model.User) (*[]model.Card, error) {
	cr := repository.NewCardRepository(cs.db)
	cards, err := cr.ByUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("cardService - CardsByUser - cr.ByUser: %w", err)
	}

	return cards, nil
}
