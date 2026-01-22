package v1handler

import (
	"net/http"

	"enterdev.com.vn/utils"
	"github.com/gin-gonic/gin"
)

var validCategory = map[string]bool{
	"golang": true,
	"php":    true,
	"python": true,
}

type CategoryHandler struct {
}

type GetCategoryByCategoryParam struct {
	Category string `uri:"category" binding:"oneof=php golang python"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	var params GetCategoryByCategoryParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": params.Category,
	})
}
