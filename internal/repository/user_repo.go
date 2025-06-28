package repository

import (
	"DucTran999/di-with-go/internal/entity"
	"context"
)

// UserRepository provides database operations for user entities.
type UserRepository struct{}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	// Simulate that insert in db success and got id 9420
	user.ID = "9420"

	return nil
}
