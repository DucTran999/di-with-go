package handler

import (
	"DucTran999/di-with-go/internal/domain/inbound"

	"github.com/google/wire"
)

var UserHandlerProvider = wire.NewSet(
	NewUserHandler,
	wire.Bind(new(inbound.UserHandler), new(*userHandler)),
)
