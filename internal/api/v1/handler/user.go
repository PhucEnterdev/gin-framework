package v1handler

import (
	"net/http"

	"enterdev.com.vn/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

type GetUserByIDV1Param struct {
	ID int `uri:"id" binding:"gt=0"`
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
	var params GetUserByIDV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandlerValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V1)",
	})
}

func (u *UserHandler) GetUserByUUIDV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")
	uid, err := utils.ValidationUUID("UUID", uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get user by UUID (V1)",
		"user_uuid": uid,
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
