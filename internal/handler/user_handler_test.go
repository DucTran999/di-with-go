package handler_test

import (
	"DucTran999/di-with-go/internal/entity"
	"DucTran999/di-with-go/internal/handler"
	"DucTran999/di-with-go/test/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	t.Run("create user failed", func(t *testing.T) {
		// Create a test Gin context
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Create a test HTTP request (e.g., POST /users with JSON body)
		c.Request, _ = http.NewRequest(http.MethodGet, "/join/daniel", nil)
		c.Params = append(c.Params, gin.Param{Key: "username", Value: "daniel"})

		ErrCreateUser := errors.New("failed to create user")
		mockUsecase := mocks.NewUserUseCase(t)
		mockUsecase.EXPECT().
			CreateUser(mock.Anything, mock.AnythingOfType("string")).
			Return(nil, ErrCreateUser)

		sut := handler.NewUserHandler(mockUsecase)

		sut.RegisterUser(c, "daniel")

		// Assert HTTP response
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("create user success", func(t *testing.T) {
		// Create a test Gin context
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Create a test HTTP request (e.g., POST /users with JSON body)
		c.Request, _ = http.NewRequest(http.MethodGet, "/join/daniel", nil)
		c.Params = append(c.Params, gin.Param{Key: "username", Value: "daniel"})

		user := &entity.User{
			Name: "daniel",
			ID:   "9420",
		}
		mockUsecase := mocks.NewUserUseCase(t)
		mockUsecase.EXPECT().
			CreateUser(mock.Anything, mock.AnythingOfType("string")).
			Return(user, nil)
		sut := handler.NewUserHandler(mockUsecase)

		sut.RegisterUser(c, "daniel")

		// Assert HTTP response
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
