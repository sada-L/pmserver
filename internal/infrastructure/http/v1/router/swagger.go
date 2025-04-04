package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func NewSwaggerRouter(g *gin.RouterGroup) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
