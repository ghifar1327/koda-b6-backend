package routes

import (
	"backend/internal/di"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, container *di.Container) {
	userhandler := container.Userhandle()

	auth := r.Group("/auth")
	{
		auth.POST("/register", userhandler.Register)
	}

}
	