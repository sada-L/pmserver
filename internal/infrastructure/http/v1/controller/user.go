package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (uc *UserController) DeleteUser(ctx *gin.Context) {
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
}
