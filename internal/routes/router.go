package routes

import (
	"backend/internal/di"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, container *di.Container) {

	userHandler := container.UserHandler()
	// Crud
	users := r.Group("/users")
	{
		users.PATCH("/:id", userHandler.UpdateUser)
	}

	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
	}

}