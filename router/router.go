package router

import (
	"github.com/Presbyter/ddd/router/handler"
	"github.com/gin-gonic/gin"
)

func RouterBind(g *gin.Engine) *gin.Engine {

	userHandler := handler.NewUserHandler(nil)
	user := g.Group("/user")
	user.POST("/login", userHandler.Login())

	return g
}
