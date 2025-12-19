package v1handler

import (
	"net/http"

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

	if !validCategory[category] {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Category must be one of: php, golang, python",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}
