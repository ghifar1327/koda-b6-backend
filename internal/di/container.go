package di

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	db *pgx.Conn
}

func NewContainer(db *pgx.Conn) *Container {
	return &Container{db: db}
}

func (c *Container) Userhandle() *handlers.UserHandler {
	userRepo := repository.NewUserrepository(c.db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserhadler(userService)

	return userHandler
}
