package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func BuildImageURL(ctx *gin.Context, path string) string {
	if path == "" {
		return ""
	}

	// get scheme (handle reverse proxy)
	scheme := ctx.GetHeader("X-Forwarded-Proto")
	if scheme == "" {
		if ctx.Request.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}

	// get host
	host := ctx.GetHeader("X-Forwarded-Host")
	if host == "" {
		host = ctx.Request.Host
	}

	// clean double slash
	cleanPath := "/" + strings.TrimPrefix(path, "/")

	return scheme + "://" + host + cleanPath
}