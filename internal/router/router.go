package router

import (
	"DucTran999/di-with-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/join/:username", func(c *gin.Context) {
		userHandler.RegisterUser(c, c.Param("username"))
	})

	return router
}
