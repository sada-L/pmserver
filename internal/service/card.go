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

func (cs *cardService) CardsByUser(ctx context.Context, user *model.User) (*[]model.Card, error) {
	cr := repository.NewCardRepository(cs.db)
	cards, err := cr.ByUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("cardService - CardsByUser - cr.ByUser: %w", err)
	}

	return cards, nil
}
