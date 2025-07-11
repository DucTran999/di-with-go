package main

import (
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/internal/repository"
	"DucTran999/di-with-go/internal/router"
	"DucTran999/di-with-go/internal/usecase"
	"DucTran999/di-with-go/internal/usecase/port"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func InitApp() *App {
	container := dig.New()

	// Provide all dependencies
	if err := container.Provide(repository.NewUserRepository, dig.As(new(port.UserRepository))); err != nil {
		log.Fatalf("failed to provide UserRepository: %v", err)
	}

	if err := container.Provide(usecase.NewUserUseCase); err != nil {
		log.Fatalf("failed to provide UserUseCase: %v", err)
	}

	if err := container.Provide(handler.NewUserHandler); err != nil {
		log.Fatalf("failed to provide UserHandler: %v", err)
	}

	if err := container.Provide(router.SetupRoutes); err != nil {
		log.Fatalf("failed to provide SetupRoutes: %v", err)
	}

	if err := container.Provide(NewApp); err != nil {
		log.Fatalf("failed to provide App: %v", err)
	}

	var app *App
	if err := container.Invoke(func(a *App) {
		app = a
	}); err != nil {
		log.Fatalf("failed to invoke App: %v", err)
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
