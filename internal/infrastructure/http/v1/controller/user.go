package controller

import (
	"errors"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
	"net/http"
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
		ctx := r.Context()
		user := utils.UserFromContext(ctx)
		if user == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if err := uc.us.DeleteUser(r.Context(), user.Id); err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusNoContent, nil)
	}
}

func (uc *UserController) UpdateUser() http.HandlerFunc {
	type Input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := utils.UserFromContext(ctx)
		if user == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if err := uc.us.UpdateUser(r.Context(), user); err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusNoContent, nil)
	}
}
