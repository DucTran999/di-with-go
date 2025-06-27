//go:build wireinject
// +build wireinject

package main

import (
	"DucTran999/di-with-go/internal/domain/inbound"
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/internal/repository"
	"DucTran999/di-with-go/internal/usecase"

	"github.com/google/wire"
)

func InitUserHandler() inbound.UserHandler {
	wire.Build(
		repository.UserRepositoryProvider,
		usecase.UserUseCaseProvider,
		handler.UserHandlerProvider,
	)
	return nil
}
