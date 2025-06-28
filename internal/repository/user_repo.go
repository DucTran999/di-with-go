package repository

import (
	"DucTran999/di-with-go/internal/entity"
	"context"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	// Simulate that insert in db success and got id 9420
	user.ID = "9420"

	return nil
}
