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

type ImageProduct struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required"`
}

type ProductAttribute struct {
	AttributeName  string `json:"attribute_name" binding:"required"`
	AttributeValue string `json:"attribute_value" binding:"required"`
}

type CreatePostV1Body struct {
	Name              string             `json:"product_name" binding:"required"`
	Price             int                `json:"price" binding:"required,min_int=10000,max_int=100000000"`
	ImageProduct      ImageProduct       `json:"image_product" binding:"required"`
	Tags              []string           `json:"tags" binding:"required,gt=0,lt=3"`
	Display           *bool              `json:"display" binding:"omitempty"`
	ProductAttributes []ProductAttribute `json:"product_attributes" binding:"required,gt=0,dive"`
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
	var body CreatePostV1Body
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	if body.Display == nil {
		defaultDisplay := true
		body.Display = &defaultDisplay
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":            "Create product (V1)",
		"product_name":       body.Name,
		"price":              body.Price,
		"image_name":         body.ImageProduct.ImageName,
		"image_link":         body.ImageProduct.ImageLink,
		"tags":               body.Tags,
		"display":            body.Display,
		"product_attributes": body.ProductAttributes,
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
