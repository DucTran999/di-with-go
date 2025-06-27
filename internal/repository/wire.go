package repository

import (
	"DucTran999/di-with-go/internal/domain"

	"github.com/google/wire"
)

var UserRepositoryProvider = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(domain.UserRepository), new(*userRepositoryImpl)),
)
