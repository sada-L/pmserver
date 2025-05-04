package controller

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
	"net/http"
	"strconv"
)

type CardController struct {
	cs model.CardService
}

func NewCardController(cs model.CardService) *CardController {
	return &CardController{cs: cs}
}

func (cc *CardController) CreateCard() http.HandlerFunc {
	type Input struct {
		Title      string `json:"title"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Website    string `json:"website"`
		Notes      string `json:"notes"`
		Image      string `json:"image"`
		IsFavorite bool   `json:"is_favorite"`
		GroupId    uint   `json:"group_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := utils.UserFromContext(ctx)

		input := Input{}
		if err := utils.ReadJSON(r.Body, &input); err != nil {
			utils.BadRequestError(w)
			return
		}

		card := &model.Card{
			Title:      input.Title,
			Username:   input.Username,
			Password:   input.Password,
			Website:    input.Website,
			Notes:      input.Notes,
			Image:      input.Image,
			IsFavorite: input.IsFavorite,
			GroupId:    input.GroupId,
			UserId:     user.Id,
		}

		id, err := cc.cs.CreateCard(r.Context(), card)
		if err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, id)
	}
}

func (cc *CardController) UpdateCard() http.HandlerFunc {
	type Input struct {
		Title      string `json:"title"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Website    string `json:"website"`
		Notes      string `json:"notes"`
		Image      string `json:"image"`
		IsFavorite bool   `json:"is_favorite"`
		GroupId    uint   `json:"group_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.BadRequestError(w)
			return
		}

		input := Input{}
		if err := utils.ReadJSON(r.Body, &input); err != nil {
			utils.BadRequestError(w)
			return
		}

		card := &model.Card{
			Id:         uint(id),
			Title:      input.Title,
			Username:   input.Username,
			Password:   input.Password,
			Website:    input.Website,
			Notes:      input.Notes,
			Image:      input.Image,
			IsFavorite: input.IsFavorite,
			GroupId:    input.GroupId,
		}

		err = cc.cs.UpdateCard(r.Context(), card)
		if err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, nil)
	}
}

func (cc *CardController) DeleteCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.BadRequestError(w)
			return
		}

		if err = cc.cs.DeleteCard(r.Context(), uint(id)); err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, nil)
	}
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
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, cards)
	}
}
