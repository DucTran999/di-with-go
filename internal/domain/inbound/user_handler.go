package inbound

import (
	"context"
)

type UserHandler interface {
	RegisterUser(ctx context.Context, username string)
}
