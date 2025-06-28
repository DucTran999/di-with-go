package repository

import (
	"DucTran999/di-with-go/internal/usecase"

	"github.com/google/wire"
)

var UserRepoProvider = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(usecase.UserRepository), new(*userRepositoryImpl)),
)
