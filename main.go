package main

import (
	"log"
	"net/http"

	v1handler "enterdev.com.vn/internal/api/v1/handler"
	"enterdev.com.vn/middleware"
	"enterdev.com.vn/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()
	r.Use(middleware.ApiKeyMiddleware(), middleware.RateLimitingMiddleware())
	go middleware.CleanUpClients()

	if err := utils.RegisterValidator(); err != nil {
		panic(err)
	}

	r.StaticFS("/images", gin.Dir("./upload", false))

	userHandler := v1handler.NewUserHandler()
	productHandler := v1handler.NewProductHandler()
	categoryHandler := v1handler.NewCategoryHandler()
	newsHandler := v1handler.NewNewsHandler()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("/", userHandler.GetUsersV1)
			user.GET("/:id", userHandler.GetUserByIDV1)
			user.GET("/admin/:uuid", userHandler.GetUserByUUIDV1)
			user.POST("/", userHandler.CreateUserV1)
			user.PUT("/:id", userHandler.UpdateUserV1)
			user.DELETE("/:id", userHandler.DeleteUserV1)
		}

		product := v1.Group("/products")
		{
			product.GET("/", productHandler.GetProductsV1)
			product.GET("/:slug", productHandler.GetProductBySlugV1)
			product.POST("/", productHandler.CreateProductV1)
			product.PUT("/:id", productHandler.UpdateProductV1)
			product.DELETE("/:id", productHandler.DeleteProductV1)
		}

		category := v1.Group("/categories")
		category.Use(middleware.SimpleMiddleware())
		{
			category.GET("/:category", categoryHandler.GetCategoryByCategoryV1)
			category.POST("/", categoryHandler.CreateCategory)
		}

		news := v1.Group("/news")
		{
			news.GET("/", newsHandler.GetNewsV1)
			news.POST("/", newsHandler.CreateNewsV1)
			news.POST("/upload-file", newsHandler.CreateUploadFileNewsV1)
			news.POST("/upload-multiple-file", newsHandler.CreateUploadMultipleFileNewsV1)
			news.GET("/:slug", middleware.ApiKeyMiddleware(), newsHandler.GetNewsV1)
		}

	}

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
