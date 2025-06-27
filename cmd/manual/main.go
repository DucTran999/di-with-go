package main

import (
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/internal/repository"
	"DucTran999/di-with-go/internal/usecase"
	"context"
)

func main() {
	userRepo := repository.NewUserRepository()
	userUC := usecase.NewUserRepository(userRepo)
	userHdl := handler.NewUserHandler(userUC)

	userHdl.RegisterUser(context.Background(), "daniel")
}
