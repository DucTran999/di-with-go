package usecase

import (
	"DucTran999/di-with-go/internal/domain"

	"github.com/google/wire"
)

var UserUseCaseProvider = wire.NewSet(
	NewUserUseCase,
	wire.Bind(new(domain.UserUseCase), new(*userUseCaseImpl)),
)
