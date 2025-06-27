package handler

import (
	"DucTran999/di-with-go/internal/domain"
	"context"
	"log"
)

type userHandler struct {
	userUC domain.UserUseCase
}

func NewUserHandler(userUC domain.UserUseCase) *userHandler {
	return &userHandler{
		userUC: userUC,
	}
}

func (hdl *userHandler) RegisterUser(ctx context.Context, username string) {
	user, err := hdl.userUC.CreateUser(ctx, username)
	if err != nil {
		log.Printf("[ERROR] failed to create user: %v", err)
		return
	}

	log.Printf("[INFO] user created name=%s, id=%s", user.Name, user.ID)
}
