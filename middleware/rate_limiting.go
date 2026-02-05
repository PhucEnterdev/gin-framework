package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	// nếu chưa từng gửi thì tạo mới
	if !exists {
		limiter := rate.NewLimiter(5, 10) // 5 request/sec , brush: số lượng tối đa có thể phục vụ cùng lúc
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient
		return limiter
	}
	// nếu đã gửi request thì cập nhật
	client.lastSeen = time.Now()
	return client.limiter
}

func CleanUpClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// ab -n 20 -c 1 -H "X-API-KEY:phuccongtu" http://localhost:8080/api/v1/categories/golang

func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many request",
				"message": "Bạn gửi quá nhiều yêu cầu. Thử lại sau.",
			})
		}
	}
}
