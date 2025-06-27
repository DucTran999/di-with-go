package domain

import "context"

type UserUseCase interface {
	CreateUser(ctx context.Context, username string) (*User, error)
}
