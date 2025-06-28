package usecase

import (
	"DucTran999/di-with-go/internal/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}

type UserUseCaseImpl struct {
	userRepo UserRepository
}

func NewUserUseCase(userRepo UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		userRepo: userRepo,
	}
}

func (r *UserUseCaseImpl) CreateUser(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{
		ID:   "",
		Name: username,
	}
	if err := r.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
