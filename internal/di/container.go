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

	//FORGOT PASSWORD
	fpRepo    *repository.ForgotPWDRepository
	fpService *service.ForgotPwdService
	fpHandler *handlers.ForgotPwdHandler

	// PRODUCT
	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handlers.ProductHandler
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

	//PRODUCT
	c.productRepo = repository.NewProductRepository(c.db)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handlers.NewProductHandler(c.productService)

	//FORGOT PASSWORD
	c.fpRepo = repository.NewForgotPWDRepository(c.db)
	c.fpService = service.NewForgotPwdService(c.userRepo, c.fpRepo)
	c.fpHandler = handlers.NewForgitPwdHandler(c.fpService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}

func (c *Container) ForgotPwdHandle() *handlers.ForgotPwdHandler {
	return c.fpHandler
}
