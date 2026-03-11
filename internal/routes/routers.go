package routes

import (
	"backend/internal/di"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, container *di.Container) {
	//=========================================================== CRUD
	//user
	userHandler := container.UserHandler()
	forgotPwdHandler := container.ForgotPwdHandle()
	productHandler := container.ProductHandler()
	reviewProductHandler := container.ReviewProductHandler()

	admin := r.Group("/admin")
	{
		users := admin.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUserById)
			users.PATCH("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
		//PRODUCT
		product := admin.Group("/products")
		{
			product.POST("", productHandler.CreateProduct)
			product.GET("", productHandler.GetProducts)
			product.GET("/:id", productHandler.GetProductbyID)
		}
		reviewProduct := admin.Group("/review-product")
		{
			reviewProduct.GET("", reviewProductHandler.GetReviewProducts)
			reviewProduct.GET("/:id", reviewProductHandler.GetReviewProductbyID)
			reviewProduct.PATCH("/:id", reviewProductHandler.UpdateReviewProduct)
			reviewProduct.DELETE("/:id", reviewProductHandler.DeleteReviewProduct)
		}
	}

	// =========================================================== FEATURE
	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("forgot-password", forgotPwdHandler.RequestForgotPwd)
		auth.PATCH("forgot-pasaword", forgotPwdHandler.Reretpassword)
	}

}
