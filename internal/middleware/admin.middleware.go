package middleware

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleId, exists := ctx.Get("role_id")
		if !exists {
			if !exists {
				ctx.JSON(http.StatusUnauthorized, dto.Response{
					Success: false,
					Message: "Unauthorized",
				})
				ctx.Abort()
				return
			}
		}
		if roleId.(int) != 1 {
			ctx.JSON(http.StatusForbidden, dto.Response{
				Success: false,
				Message: "Forbidden: admin only",
			})
			ctx.Abort()
			return
		}
	}
}
