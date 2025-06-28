package main

import (
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/internal/repository"
	"DucTran999/di-with-go/internal/router"
	"DucTran999/di-with-go/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func InitApp() *App {
	container := dig.New()

	// Provide all dependencies
	container.Provide(repository.NewUserRepository, dig.As(new(usecase.UserRepository)))
	container.Provide(usecase.NewUserUseCase, dig.As(new(handler.UserUseCase)))
	container.Provide(handler.NewUserHandler, dig.As(new(router.UserHandler)))
	container.Provide(router.SetupRoutes)
	container.Provide(NewApp)

	var app *App

	err := container.Invoke(func(a *App) {
		app = a
	})
	if err != nil {
		log.Fatalf("failed to build App: %v", err)
	}

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
