package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Trước khi bắt đầu vào handler
		log.Println("Start func check from middleware")

		ctx.Writer.Write([]byte("start middleware\n"))
		ctx.Next()

		// Sau khi handler xử lý xong
		log.Println("End func check from middleware")
		ctx.Writer.Write([]byte("\nend middleware\n"))

	}
}
