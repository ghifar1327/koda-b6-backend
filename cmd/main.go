package main

import (
	"backend/internal/di"
	"backend/internal/routes"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)


// @title Coffee Shop API
// @version 1.0
// @description This is a coffee shop backend API
// @host localhost:8888
// @BasePath /
func main() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	container := di.NewContainer(conn)

	routes.Router(r, container)

	r.Run(":8888")
}