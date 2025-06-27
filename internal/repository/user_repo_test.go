package repository_test

import (
	"DucTran999/di-with-go/internal/domain"
	"DucTran999/di-with-go/internal/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	sut := repository.NewUserRepository()
	user := &domain.User{
		Name: "daniel",
	}
	expectedID := "9420"

	err := sut.Create(context.Background(), user)

	require.NoError(t, err)
	assert.Equal(t, expectedID, user.ID)
}
