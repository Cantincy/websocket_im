package router

import (
	"github.com/gin-gonic/gin"
	"newim/handler"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()

	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", handler.UserRegisterHandler)
		userGroup.GET("/ws", handler.UserWebSocketHandler)
	}

	return engine
}
