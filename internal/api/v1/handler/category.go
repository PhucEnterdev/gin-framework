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

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	category := ctx.Param("category")

	if err := utils.ValidationInList("Category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}
