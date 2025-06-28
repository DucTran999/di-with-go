package router

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(ctx *gin.Context, username string)
}

func SetupRoutes(userHandler UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/join/:username", func(c *gin.Context) {
		userHandler.RegisterUser(c, c.Param("username"))
	})

	return router
}
