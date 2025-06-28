package usecase

import (
	"DucTran999/di-with-go/internal/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
}

// UserUseCase handles business logic for user operations.
type UserUseCase struct {
	userRepo UserRepository
}

// NewUserUseCase creates a new instance of UserUseCase.
func NewUserUseCase(userRepo UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (r *UserUseCase) CreateUser(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{
		ID:   "",
		Name: username,
	}
	if err := r.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
