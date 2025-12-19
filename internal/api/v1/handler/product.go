package v1handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
	})
}

func (p *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Slug must contain only lowercase letter, numbers, hyphens and dots",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by Slug (V1)",
		"slug":    slug,
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
