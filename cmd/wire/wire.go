//go:build wireinject
// +build wireinject

package main

import (
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/internal/repository"
	"DucTran999/di-with-go/internal/router"
	"DucTran999/di-with-go/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

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

func InitApp() *App {
	wire.Build(
		repository.NewUserRepository,
		wire.Bind(new(usecase.UserRepository), new(*repository.UserRepository)),
		usecase.NewUserUseCase,
		wire.Bind(new(handler.UserUseCase), new(*usecase.UserUseCase)),
		handler.NewUserHandler,
		wire.Bind(new(router.UserHandler), new(*handler.UserHandler)),
		router.SetupRoutes,
		NewApp,
	)
	return &App{}
}
