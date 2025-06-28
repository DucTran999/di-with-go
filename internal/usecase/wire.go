package usecase

import (
	"DucTran999/di-with-go/internal/handler"

	"github.com/google/wire"
)

// wire want to access private userUseCaseImpl struct
// and import interface form handler is Dependency Violation.
// Another options is public userUseCaseImpl.
var UserUsecaseProvider = wire.NewSet(
	NewUserUseCase,
	wire.Bind(new(handler.UserUseCase), new(*userUseCaseImpl)),
)
