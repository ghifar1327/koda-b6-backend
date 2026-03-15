package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend/docs"
	"backend/internal/di"
)

func Router(r *gin.Engine, container *di.Container) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//user
	userHandler := container.UserHandler()
	forgotPwdHandler := container.ForgotPwdHandle()
	productHandler := container.ProductHandler()
	landingHandler := container.LandingHandler()
	reviewProductHandler := container.ReviewProductHandler()
	transactionHandler := container.TransactionHandler()

	admin := r.Group("/admin")
	{
		//users
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

	}
	// =========================================================== FEATURE

	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/forgot-password", forgotPwdHandler.RequestForgotPwd)
		auth.PATCH("/reset-password", forgotPwdHandler.ResetPassword)
	}

	// TRANSACSION
	transaction := r.Group("/transaction")
	{
		transaction.GET("", transactionHandler.GetAllTransactions)
		transaction.GET("/:id", transactionHandler.GetTransactionbyID)
	}

	//reviews
	reviewProduct := r.Group("/review-product")
	{
		reviewProduct.GET("", reviewProductHandler.GetAllReviewProducts)
		reviewProduct.GET("/:id", reviewProductHandler.GetReviewProductbyID)
		reviewProduct.PATCH("/:id", reviewProductHandler.UpdateReviewProduct)
		reviewProduct.DELETE("/:id", reviewProductHandler.DeleteReviewProduct)
	}

	// Landing
	landing := r.Group("/")
	{
		landing.GET("/reviews", landingHandler.GetAllReviewProducts)
		landing.GET("/reviews/:id", landingHandler.GetRecommendedProductByID)
		landing.GET("/recommended-product", landingHandler.GetRecommendedProducts)
		landing.GET("/recommended-product/:id", landingHandler.GetRecommendedProductByID)
	}
}
