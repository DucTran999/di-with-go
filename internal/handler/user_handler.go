package handler

import (
	"DucTran999/di-with-go/internal/usecase/port"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// userHandler handles HTTP requests for user operations.
type UserHandler struct {
	userUC port.UserUsecase
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userUC port.UserUsecase) *UserHandler {
	return &UserHandler{
		userUC: userUC,
	}
}

func (hdl *UserHandler) RegisterUser(ctx *gin.Context, username string) {
	user, err := hdl.userUC.CreateUser(ctx, username)
	if err != nil {
		log.Printf("[ERROR] failed to create user: %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	successResp := map[string]any{
		"code":    200,
		"message": fmt.Sprintf("%s(%s): playing di", user.Name, user.ID),
	}

	ctx.JSON(http.StatusOK, successResp)
}
