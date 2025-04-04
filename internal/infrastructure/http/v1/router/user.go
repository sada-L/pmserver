package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewUserRouter(uc *controller.UserController, g *gin.RouterGroup) {
	g.POST("/login", uc.Login)
	g.POST("/signup", uc.Signup)
}
