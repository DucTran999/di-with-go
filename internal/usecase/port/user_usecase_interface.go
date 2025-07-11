package port

import (
	"DucTran999/di-with-go/internal/entity"
	"context"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, username string) (*entity.User, error)
}
