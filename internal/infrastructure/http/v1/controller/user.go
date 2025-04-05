package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
)

type UserController struct {
	us model.UserService
}

func NewUserController(us model.UserService) *UserController {
	return &UserController{us: us}
}

var validate *validator.Validate

func (uc *UserController) LoginUser() http.HandlerFunc {
	type Input struct {
		User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}

		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		user, err := uc.us.Authenticate(r.Context(), input.User.Email, input.User.Password)
		if err != nil || user == nil {
			utils.InvalidUserCredentialsError(w)
			return
		}

		token, err := utils.GenerateUserToken(user)
		if err != nil {
			utils.ServerError(w, err)
			return
		}

		user.Token = token

		utils.WriteJSON(w, http.StatusOK, utils.M{"user": user})
	}
}

func (uc *UserController) CreateUser() http.HandlerFunc {
	type Input struct {
		User struct {
			Email    string `json:"email" validate:"required,email"`
			Username string `json:"username" validate:"required,min=2"`
			Password string `json:"password" validate:"required,min=8,max=72"`
		} `json:"user" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		input := &Input{}

		if err := utils.ReadJSON(r.Body, input); err != nil {
			utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		}

		// if err := validate.Struct(input.User); err != nil {
		// 	utils.ValidationError(w, err)
		// 	return
		// }

		user := model.User{
			Email:    input.User.Email,
			Username: input.User.Username,
		}

		user.SetPassword(input.User.Password)

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

		utils.WriteJSON(w, http.StatusCreated, utils.M{"user": user})
	}
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
}
