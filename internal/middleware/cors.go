package middleware

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	var origins []string
	_ = json.Unmarshal([]byte(os.Getenv("CORS_ORIGINS")), &origins)

	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")

		for _, o := range origins {
			if o == origin {
				ctx.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}

		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}

		ctx.Next()
	}
}