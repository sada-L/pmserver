package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sada-L/pmserver/internal/model"
)

type UserController struct {
	UserService model.UserService
}

func NewUserController(s model.UserService) *UserController {
	return &UserController{UserService: s}
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Login
//
//	@Summary		Авторизация.
//	@Description	Авторизация по логину и паролю.
//	@Tags			Auth
//	@Produce		json
//	@Param			email		formData	string	true	"Email"
//	@Param			password	formData	string	true	"Password"
//	@Success		200			{object}	model.User
//	@Failure		401			{object}	string
//	@Failure		500			{object}	string
//	@Router			/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var request LoginRequest

	err := ctx.Bind(&request)
	if err != nil {
		log.Print(fmt.Errorf("failed to bind request: %w", err))
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err})
		return
	}

	user, err := c.UserService.UserByEmail(ctx, request.Email)
	if err != nil {
		log.Print(fmt.Errorf("user not found: %w", err))
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type SignUpRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Signup
//
//	@Summary		Signup.
//	@Description	Signup
//	@Tags			Auth
//	@Produce		json
//	@Param			name	formData	string	true	"Name"
//	@Param			email		formData	string	true	"Email"
//	@Param			password	formData	string	true	"Password"
//	@Success		200			{object}	model.User
//	@Failure		401			{object}	string
//	@Failure		500			{object}	string
//	@Router			/signup [post]
func (c *UserController) Signup(ctx *gin.Context) {
	var request SignUpRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		log.Print(fmt.Errorf("failed to bind request: %w", err))
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err})
		return
	}

	user := &model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = c.UserService.CreateUser(ctx, user)
	if err != nil {
		log.Print(fmt.Errorf("server error: %w", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": err})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
}
