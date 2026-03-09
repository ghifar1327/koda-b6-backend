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
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUserById)
		users.PATCH("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
	}

}