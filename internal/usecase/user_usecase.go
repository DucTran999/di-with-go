package usecase

import (
	"DucTran999/di-with-go/internal/entity"
	"DucTran999/di-with-go/internal/usecase/port"
	"context"
)

// UserUsecase handles business logic for user operations.
type userUsecase struct {
	userRepo port.UserRepository
}

// NewUserUsecase creates a new instance of UserUsecase.
func NewUserUseCase(userRepo port.UserRepository) port.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (r *userUsecase) CreateUser(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{
		ID:   "",
		Name: username,
	}
	if err := r.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
