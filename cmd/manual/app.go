package main

import (
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/internal/repository"
	"DucTran999/di-with-go/internal/router"
	"DucTran999/di-with-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

func InitApp() *App {
	// DI manual
	userRepo := repository.NewUserRepository()
	userUC := usecase.NewUserUseCase(userRepo)
	userHdl := handler.NewUserHandler(userUC)
	r := router.SetupRoutes(userHdl)
	app := NewApp(r)

	return app
}

type App struct {
	router *gin.Engine
}

func NewApp(router *gin.Engine) *App {
	return &App{
		router: router,
	}
}

func (a *App) Run(address string) error {
	return a.router.Run(address)
}
