package usecase

import (
	"DucTran999/di-with-go/internal/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}

type userUseCaseImpl struct {
	userRepo UserRepository
}

func NewUserUseCase(userRepo UserRepository) *userUseCaseImpl {
	return &userUseCaseImpl{
		userRepo: userRepo,
	}
}

func (r *userUseCaseImpl) CreateUser(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{
		Name: username,
	}
	if err := r.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
