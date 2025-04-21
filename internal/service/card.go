package service

import (
	"context"
	"github.com/sada-L/pmserver/internal/model"
)

type cardService struct {
	cr *model.CardRepository
}

func NewCardService(cr *model.CardRepository) model.CardService {
	return &cardService{cr: cr}
}

func (cs *cardService) CardsByUserId(ctx context.Context, id string) (*[]model.Card, error) {
	return nil, nil
}
