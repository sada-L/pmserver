package controller

import (
	"errors"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
	"net/http"
	"strconv"
)

type AuthController struct {
	us model.UserService
}

func NewAuthController(us model.UserService) *AuthController {
	return &AuthController{us: us}
}

func (ac *AuthController) Login() http.HandlerFunc {
	type Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}
		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.BadRequestError(w)
			return
		}

		user, err := ac.us.UserByEmail(r.Context(), input.Email)
		if err != nil || user == nil {
			utils.InvalidUserCredentialsError(w)
			return
		}

		if !user.VerifyPassword(input.Password) {
			utils.InvalidUserCredentialsError(w)
			return
		}

		authResponse, err := utils.NewAuthResponse(user, 5)
		if err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, authResponse)
	}
}

func (ac *AuthController) Register() http.HandlerFunc {
	type Input struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,min=2"`
		Password string `json:"password" validate:"required,min=8,max=72"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}
		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.BadRequestError(w)
			return
		}

		user := &model.User{
			Email:    input.Email,
			Username: input.Username,
		}

		if err := user.SetPassword(input.Password); err != nil {
			utils.BadRequestError(w)
			return
		}

		if err := ac.us.CreateUser(r.Context(), user); err != nil {
			switch {
			case errors.Is(err, model.ErrDuplicateEmail):
				err = model.ErrorM{"email": {"this email is already in use"}}
				utils.ErrorResponse(w, http.StatusConflict, err)

			case errors.Is(err, model.ErrDuplicateUsername):
				err = model.ErrorM{"username": {"this username is already in use"}}
				utils.ErrorResponse(w, http.StatusConflict, err)
			default:
				utils.InternalError(w, err)
			}
			return
		}

		authResponse, err := utils.NewAuthResponse(user, 5)
		if err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, authResponse)
	}
}

func (ac *AuthController) RefreshToken() http.HandlerFunc {
	type Input struct {
		RefreshToken string `json:"refresh_token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}
		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.BadRequestError(w)
			return
		}

		claims, err := utils.ParseUserToken(input.RefreshToken)
		if err != nil {
			utils.InvalidAuthTokenError(w)
			return
		}

		id := claims["id"].(float64)
		sid := strconv.FormatFloat(id, 'f', -1, 64)
		user, err := ac.us.UserById(r.Context(), sid)
		if err != nil || user == nil {
			utils.InvalidUserCredentialsError(w)
			return
		}

		authResponse, err := utils.NewAuthResponse(user, 5)
		if err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, authResponse)
	}
}
