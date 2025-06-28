package handler

import (
	"DucTran999/di-with-go/internal/router"

	"github.com/google/wire"
)

var UserHandlerProvider = wire.NewSet(
	NewUserHandler,
	wire.Bind(new(router.UserHandler), new(*userHandler)),
)
