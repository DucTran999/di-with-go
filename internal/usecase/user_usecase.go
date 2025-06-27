package usecase

import (
	"DucTran999/di-with-go/internal/domain"
	"context"
)

type userUseCaseImpl struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *userUseCaseImpl {
	return &userUseCaseImpl{
		userRepo: userRepo,
	}
}

func (r *userUseCaseImpl) CreateUser(ctx context.Context, username string) (*domain.User, error) {
	user := domain.User{
		Name: username,
	}
	if err := r.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
