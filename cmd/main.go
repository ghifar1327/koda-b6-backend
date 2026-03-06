package main

import (
	"backend/internal/di"
	"backend/internal/routes"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	dbURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	container := di.NewContainer(conn)

	routes.Router(r, container)

	r.Run(":8080")
}

// import (
// 	"backend/internal/handlers"
// 	"backend/internal/middleware"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	godotenv.Load()
// 	// ===================================================================== CRUD USER
// 	r := gin.Default()
// 	r.Use(middleware.CorsMiddleware())
// 	users := r.Group("/users")
// 	users.POST("", handlers.CreateUser)
// 	users.GET("", handlers.GetAllUser)
// 	users.GET("/:id", handlers.GetUserByID)
// 	users.PATCH("/:id", handlers.UpdateUser)
// 	users.DELETE("/:id", handlers.DeleteUser)

// 	// ================================================================ CRUD PRODUCT
// 	products := r.Group("/products")
// 	products.GET("/products", handlers.GetAllProduct)

// 	// ===================================================================== Feature auth
// 	r.POST("/login", handlers.AuthLogin)

// 	// ====================================================================== FEATURE ORDER
// 	orders := r.Group("/orders")
// 	orders.POST("/addcart", handlers.AddChart)
// 	orders.POST("/checkout", handlers.Checkout)

// 	r.Run("localhost:8888")
// }
