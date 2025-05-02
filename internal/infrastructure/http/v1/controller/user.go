package controller

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
)

type UserController struct {
	us model.UserService
}

func NewUserController(us model.UserService) *UserController {
	return &UserController{us: us}
}

func (uc *UserController) GetCurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := utils.UserFromContext(ctx)
		if user == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		utils.WriteJSON(w, http.StatusOK, user)
	}
}

func (uc *UserController) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if err := uc.us.DeleteUser(r.Context(), uint(id)); err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusNoContent, nil)
	}
}

func (uc *UserController) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
