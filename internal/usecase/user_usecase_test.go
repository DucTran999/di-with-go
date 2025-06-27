package usecase_test

import (
	"DucTran999/di-with-go/internal/domain"
	"DucTran999/di-with-go/internal/usecase"
	"DucTran999/di-with-go/test/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	t.Run("insert in db failed should return error", func(t *testing.T) {
		// Arrange
		ErrCreateUser := errors.New("failed to insert into DB")
		mockRepo := mocks.NewUserRepository(t)
		mockRepo.EXPECT().
			Create(mock.Anything, mock.AnythingOfType("*domain.User")).
			Return(ErrCreateUser)
		sut := usecase.NewUserUseCase(mockRepo)

		// Act
		user, err := sut.CreateUser(context.Background(), "daniel")

		// Assert
		require.ErrorIs(t, err, ErrCreateUser)
		assert.Nil(t, user)
	})

	t.Run("insert success", func(t *testing.T) {
		// Arrange
		mockRepo := mocks.NewUserRepository(t)
		mockRepo.EXPECT().
			Create(mock.Anything, mock.AnythingOfType("*domain.User")).
			Run(func(ctx context.Context, user *domain.User) {
				user.ID = "8888"
			}).
			Return(nil)
		sut := usecase.NewUserUseCase(mockRepo)

		// Act
		user, err := sut.CreateUser(context.Background(), "daniel")

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "8888", user.ID)
	})
}
