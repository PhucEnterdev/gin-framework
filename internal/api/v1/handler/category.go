package v1handler

import (
	"log"
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
type CreateCategoryV1Body struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
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

	log.Println("Into GetCategoryByCategoryV1 ")

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": params.Category,
	})
}

func (c *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var body CreateCategoryV1Body
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create category successfully",
		"name":    body.Name,
		"status":  body.Status,
	})
}
