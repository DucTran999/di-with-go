package domain

import "context"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
}
