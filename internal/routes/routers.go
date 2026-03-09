package routes

import (
	"backend/internal/di"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, container *di.Container) {
	//=========================================================== CRUD
	//user
	userHandler := container.UserHandler()
	users := r.Group("/users")
	{
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUserById)
		users.PATCH("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
	//PRODUCT
	productHandler := container.ProductHandler()
	product := r.Group("/products")
	{
		product.GET("", productHandler.GetProducts)
		product.GET("/:id", userHandler.GetUserById)
	}



	// =========================================================== FEATURE
	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
	}

}
