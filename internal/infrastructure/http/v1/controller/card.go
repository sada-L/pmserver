package controller

import (
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
	"net/http"
)

type CardController struct {
	cs model.CardService
}

func NewCardController(cs model.CardService) *CardController {
	return &CardController{cs: cs}
}

func (cc *CardController) GetCardsByUserId() http.HandlerFunc {
	type Input struct {
		UserId string `json:"user_id" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := Input{}

		if err := utils.ReadJSON(r.Body, &input); err != nil {
			utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		}

		cards, err := cc.cs.CardsByUserId(r.Context(), input.UserId)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, cards)
	}
}
