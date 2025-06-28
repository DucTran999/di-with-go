package main

import (
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/internal/repository"
	"DucTran999/di-with-go/internal/router"
	"DucTran999/di-with-go/internal/usecase"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func InitApp() *App {
	var app *App

	fxApp := fx.New(
		fx.Provide(
			fx.Annotate(
				repository.NewUserRepository,
				fx.As(new(usecase.UserRepository)),
			),
			fx.Annotate(
				usecase.NewUserUseCase,
				fx.As(new(handler.UserUseCase)),
			),
			fx.Annotate(
				handler.NewUserHandler,
				fx.As(new(router.UserHandler)),
			),
			router.SetupRoutes,
			NewApp,
		),
		fx.Invoke(func(a *App) {
			app = a
		}),
	)

	if err := fxApp.Start(context.Background()); err != nil {
		log.Fatalf("failed to start fx app: %v", err)
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
