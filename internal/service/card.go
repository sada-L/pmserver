package service

import "github.com/sada-L/pmserver/internal/model"

type cardService struct {
  CardRepository *model.CardRepository  
}

func NewCardService(r *model.CardRepository) model.CardService  {
  return &cardService{CardRepository: r} 
}

func (s cardService) Create()  {
  
}
