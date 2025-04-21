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

func (uc *UserController) LoginUser() http.HandlerFunc {
	type Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}

		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		user, err := uc.us.Authenticate(r.Context(), input.Email, input.Password)
		if err != nil || user == nil {
			utils.InvalidUserCredentialsError(w)
			return
		}

		token, err := utils.GenerateUserToken(user)
		if err != nil {
			utils.ServerError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, token)
	}
}

func (uc *UserController) CreateUser() http.HandlerFunc {
	type Input struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,min=2"`
		Password string `json:"password" validate:"required,min=8,max=72"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}

		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		}

		user := model.User{
			Email:    input.Email,
			Username: input.Username,
		}

		if err := user.SetPassword(input.Password); err != nil {
			utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		}

		if err := uc.us.CreateUser(r.Context(), &user); err != nil {
			switch {
			case errors.Is(err, model.ErrDuplicateEmail):
				err = model.ErrorM{"email": {"this email is already in use"}}
				utils.ErrorResponse(w, http.StatusConflict, err)

			case errors.Is(err, model.ErrDuplicateUsername):
				err = model.ErrorM{"username": {"this username is already in use"}}
				utils.ErrorResponse(w, http.StatusConflict, err)
			default:
				utils.ServerError(w, err)
			}
			return
		}

		token, err := utils.GenerateUserToken(&user)
		if err != nil {
			utils.ServerError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, token)
	}
}

func (uc *UserController) GetUserByEmail() http.HandlerFunc {
	type Input struct {
		Email string `json:"email" validate:"required,email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}

		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		}

		user, err := uc.us.UserByEmail(r.Context(), input.Email)
		if err != nil || user == nil {
			utils.InvalidUserCredentialsError(w)
			return
		}

		utils.WriteJSON(w, http.StatusOK, user)
	}
}
func (uc *UserController) DeleteUser(ctx *gin.Context) {
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
}
