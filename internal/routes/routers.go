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

	r.Use(middleware.CorsMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "welcome to backend",
		})
	})

	// handler
	userHandler := container.UserHandler()
	authHandler := container.AuthHandler()
	productHandler := container.ProductHandler()
	cartHandelr := container.CartHandler()
	landingHandler := container.LandingHandler()
	reviewHandler := container.ReviewProductHandler()
	transactionHandler := container.TransactionHandler()
	masterHandler := container.MasterHandler()

	// ======================================================================== PUBLIC

	product := r.Group("/products")
	{
		product.GET("", productHandler.GetProducts)
		product.GET("/:id", productHandler.GetProductbyID)
		product.GET("/:id/variants", productHandler.GetVariantByIdProduct)
		product.GET("/:id/sizes", productHandler.GetSizesByIdProduct)
	}

	
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/forgot-password", authHandler.RequestForgotPwd)
		auth.PATCH("/reset-password", authHandler.ResetPassword)
	}
	
	review := r.Group("/review-product")
	{
		review.GET("", reviewHandler.GetAllReviewProducts)
		review.GET("/:id", reviewHandler.GetReviewProductbyID)
	}
	
	landing := r.Group("/landing")
	{
		landing.GET("/reviews", landingHandler.GetAllReviewProductsLanding)
		landing.GET("/reviews/:id", landingHandler.GetRecommendedProductByID)
		landing.GET("/recommended-product", landingHandler.GetRecommendedProducts)
		landing.GET("/recommended-product/:id", landingHandler.GetRecommendedProductByID)
	}
	
	master := r.Group("/master")
	{
		master.GET("/:table", masterHandler.GetAll)
		master.GET("/:table/:id", masterHandler.GetById)
	}
	
	// ========================================================================= USER (LOGIN REQUIRED)
	
	user := r.Group("/")
	// user.Use(middleware.AuthMiddleware())
	{
		transaction := user.Group("/transactions")
		{
			transaction.POST("", transactionHandler.CreateTransaction)
			transaction.GET("", transactionHandler.GetAllTransactions)
			transaction.GET("/:id", transactionHandler.GetTransactionbyID)
			transaction.GET("/user/:id", transactionHandler.GetAllTransactionsByUserID)

		}
		cart := user.Group("/cart")
		{
			cart.GET("/:user_id", cartHandelr.GetCartByUserId)
			cart.POST("",cartHandelr.AddCart)
			cart.DELETE("/:id", cartHandelr.DeleteCart)
		}
		
		review := user.Group("/review-product")
		{
			review.POST("", reviewHandler.CreateReviewProduct)
			review.PATCH("/:id", reviewHandler.UpdateReviewProduct)
			review.DELETE("/:id", reviewHandler.DeleteReviewProduct)
		}
		auth := r.Group("/auth")
		{
			auth.PATCH("/:id/update", authHandler.UpdateProfile)
		}
	}

	// ========================================================================== ADMIN

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		users := admin.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUserById)
			users.PATCH("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		transaction := admin.Group("/transaction")
		{
			transaction.PATCH("/:id", transactionHandler.UpdateTransaction)
			transaction.DELETE("/:id", transactionHandler.DeleteTransaction)
		}

		master := admin.Group("/master")
		{
			master.POST("/:table", masterHandler.Create)
			master.PATCH("/:table/:id", masterHandler.Update)
			master.DELETE("/:table/:id", masterHandler.Delete)
		}
	}
}
