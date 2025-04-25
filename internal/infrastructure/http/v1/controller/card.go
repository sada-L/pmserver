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

func (cc *CardController) GetCardsByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := utils.UserFromContext(ctx)
		if user == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		cards, err := cc.cs.CardsByUser(r.Context(), user)
		if err != nil {
			utils.ServerError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, cards)
	}
}
