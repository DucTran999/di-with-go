package handler_test

import (
	"DucTran999/di-with-go/internal/domain"
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/test/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	t.Run("create user failed", func(t *testing.T) {
		ErrCreateUser := errors.New("failed to create user")
		mockUsecase := mocks.NewUserUseCase(t)
		mockUsecase.EXPECT().
			CreateUser(mock.Anything, mock.AnythingOfType("string")).
			Return(nil, ErrCreateUser)
		sut := handler.NewUserHandler(mockUsecase)

		sut.RegisterUser(context.Background(), "daniel")
	})

	t.Run("create user success", func(t *testing.T) {
		user := &domain.User{
			Name: "daniel",
			ID:   "9420",
		}
		mockUsecase := mocks.NewUserUseCase(t)
		mockUsecase.EXPECT().
			CreateUser(mock.Anything, mock.AnythingOfType("string")).
			Return(user, nil)
		sut := handler.NewUserHandler(mockUsecase)

		sut.RegisterUser(context.Background(), "daniel")
	})
}
