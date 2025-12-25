package v1handler

import (
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

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
		"limit":   limit,
	})
}

func (p *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if err := utils.ValidationRegex(slug, slugRegex, "Slug must contain only lowercase letter, numbers, hyphens and dots"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
