package di

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	db *pgx.Conn

	// USER
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler
}

func NewContainer(db *pgx.Conn) *Container {
	container := &Container{
		db: db,
	}

	container.initDependencies()

	return container
}

func (c *Container) initDependencies() {

	// USER
	c.userRepo = repository.NewUserrepository(c.db)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserhadler(c.userService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}