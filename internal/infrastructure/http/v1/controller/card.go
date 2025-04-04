package controller

import "github.com/sada-L/pmserver/internal/model"

type cardController struct{
  CardService *model.CardService 
} 

func NewCardController(s *model.CardService) model.CardController  {
  return &cardController{CardService: s} 
}

func (c cardController) Create() {
  
}
