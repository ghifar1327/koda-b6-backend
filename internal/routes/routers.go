package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend/docs"
	"backend/internal/di"
	"backend/internal/middleware"
)

func Router(r *gin.Engine, container *di.Container) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.CorsMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"succes": "wellcome to backend",
		})
	})
	//user
	userHandler := container.UserHandler()
	AuthHandler := container.AuthHandler()
	productHandler := container.ProductHandler()
	landingHandler := container.LandingHandler()
	reviewProductHandler := container.ReviewProductHandler()
	transactionHandler := container.TransactionHandler()

	admin := r.Group("/admin")
	// admin.Use(middleware.AuthMiddleware())
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
		
	}
	// =========================================================== FEATURE
	
	product := r.Group("/products")
	{
		product.POST("", productHandler.CreateProduct)
		product.GET("", productHandler.GetProducts)
		product.GET("/:id", productHandler.GetProductbyID)
		product.GET("/:id/variants", productHandler.GetVariantByIdProduct)
		product.GET("/:id/sizes", productHandler.GetSizesByIdProduct)
	}
	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", AuthHandler.Register)
		auth.POST("/login", AuthHandler.Login)
		auth.POST("/forgot-password", AuthHandler.RequestForgotPwd)
		auth.PATCH("/reset-password", AuthHandler.ResetPassword)
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
		reviewProduct.POST("", reviewProductHandler.CreateReviewProduct)
		reviewProduct.GET("", reviewProductHandler.GetAllReviewProducts)
		reviewProduct.GET("/:id", reviewProductHandler.GetReviewProductbyID)
		reviewProduct.PATCH("/:id", reviewProductHandler.UpdateReviewProduct)
		reviewProduct.DELETE("/:id", reviewProductHandler.DeleteReviewProduct)
	}

	// Landing
	landing := r.Group("landing")
	{
		landing.GET("/reviews", landingHandler.GetAllReviewProducts)
		landing.GET("/reviews/:id", landingHandler.GetRecommendedProductByID)
		landing.GET("/recommended-product", landingHandler.GetRecommendedProducts)
		landing.GET("/recommended-product/:id", landingHandler.GetRecommendedProductByID)
	}
}
