package port

import (
	"DucTran999/di-with-go/internal/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}
