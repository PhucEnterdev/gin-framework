package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all users (V1)",
	})
}

func (u *UserHandler) GetUserByIDV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V1)",
	})
}

func (u *UserHandler) GetUserByUUIDV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")
	_, err := uuid.Parse(uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID must be a valid UUID",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get user by UUID (V1)",
		"user_uuid": uuidStr,
	})
}

func (u *UserHandler) CreateUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create user (V1)",
	})
}

func (u *UserHandler) UpdateUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user (V1)",
	})
}

func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
