package v1handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"enterdev.com.vn/utils"
	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

type CreateNewsV1Body struct {
	Titile string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required"`
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (n *NewsHandler) GetNewsV1(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get news (V1)",
			"slug":    "Không có tin tức",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get news (V1)",
			"slug":    slug,
		})
	}
}

func (n *NewsHandler) CreateNewsV1(ctx *gin.Context) {
	var body CreateNewsV1Body
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "file is required",
		})
		return
	}

	// limit file size < 5MB
	// 1 << 20 = 1 * 2^20 = 1048576 = 1MB
	// 5 << 20 = 5 * 2^20 = 1048576 * 5 = 5MB
	if image.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "file too large (5MB)",
		})
		return
	}

	// os.ModePerm = 0777 read write execute for all
	err = os.MkdirAll("./upload", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot create upload folder",
		})
		return
	}
	des := fmt.Sprintf("./upload/%s", filepath.Base(image.Filename))
	if err := ctx.SaveUploadedFile(image, des); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot save file",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create news successfully",
		"title":   body.Titile,
		"status":  body.Status,
		"image":   image.Filename,
		"path":    des,
	})

}

func (n *NewsHandler) CreateUploadFileNewsV1(ctx *gin.Context) {
	var body CreateNewsV1Body
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "file is required",
		})
		return
	}

	filename, err := utils.ValidateAndSaveFile(image, "./upload")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create news successfully",
		"title":   body.Titile,
		"status":  body.Status,
		"image":   filename,
		"path":    "./upload/" + filename,
	})
}

func (n *NewsHandler) CreateUploadMultipleFileNewsV1(ctx *gin.Context) {
	var body CreateNewsV1Body
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid multipart form",
		})
		return
	}

	images := form.File["images"]
	if len(images) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "no file provided",
		})
		return
	}

	var successFiles []string
	var failedFiles []map[string]string
	for _, image := range images {
		filename, err := utils.ValidateAndSaveFile(image, "./upload")
		if err != nil {
			failedFiles = append(failedFiles, map[string]string{
				"filename": image.Filename,
				"error":    err.Error(),
			})
			continue
		}

		successFiles = append(successFiles, filename)
	}

	resp := gin.H{
		"message":       "Create news successfully",
		"title":         body.Titile,
		"status":        body.Status,
		"success_files": successFiles,
	}

	if len(failedFiles) > 0 {
		resp["message"] = "Upload file completed with partical errors"
		resp["failed_files"] = failedFiles
	}

	ctx.JSON(http.StatusCreated, resp)
}
