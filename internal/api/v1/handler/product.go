package v1handler

import (
	"log"
	"net/http"
	"regexp"

	"enterdev.com.vn/utils"
	"github.com/gin-gonic/gin"
)

var (
	slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
)

type ProductHandler struct {
}

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=2,max=50,search"`
}

type GetProductBySlugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=2,max=5"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	var params GetProductsV1Param
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
		"search":  params.Search,
	})
}

func (p *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	var params GetProductBySlugV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by Slug (V1)",
		"slug":    params.Slug,
	})
}

func (p *ProductHandler) CreateProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create product (V1)",
	})
}

func (p *ProductHandler) UpdateProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update product (V1)",
	})
}

func (p *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
