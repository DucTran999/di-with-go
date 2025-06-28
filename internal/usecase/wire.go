package usecase

import (
	"DucTran999/di-with-go/internal/handler"

	"github.com/google/wire"
)

// Another options is public userUseCaseImpl
// import interface form handler is Dependency Violation
var UserUsecaseProvider = wire.NewSet(
	NewUserUseCase,
	wire.Bind(new(handler.UserUseCase), new(*userUseCaseImpl)),
)
