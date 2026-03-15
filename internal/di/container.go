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

	// LANDING
	landingRepo    *repository.LandingRepository
	landingService *service.LandingService
	landingHandler *handlers.LandingHandler

	//TRANSACTION
	transactionRepo    *repository.TransactionRepository
	transactionService *service.TransactionService
	transactionHandler *handlers.TransactionHandler

	// PRODUCT
	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handlers.ProductHandler

	// REVIEW PRODUCT
	reviewProductRepo    *repository.ReviewProductRepository
	reviewProductSevice  *service.ReviewProductService
	reviewProductHandler *handlers.ReviewProductHandler
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
	c.userHandler = handlers.NewUserHandler(c.userService)

	//FORGOT PASSWORD
	c.fpRepo = repository.NewForgotPWDRepository(c.db)
	c.fpService = service.NewForgotPwdService(c.userRepo, c.fpRepo)
	c.fpHandler = handlers.NewForgotPwdHandler(c.fpService)

	//LANDING
	c.landingRepo = repository.NewLandingRepository(c.db)
	c.landingService = service.NewLandingService(c.landingRepo)
	c.landingHandler = handlers.NewLandingHandler(c.landingService)

	// TRAMSACTION
	c.transactionRepo = repository.NewTransactionRepository(c.db)
	c.transactionService = service.NewTransactionService(c.transactionRepo)
	c.transactionHandler = handlers.NewTransactionHandler(c.transactionService)

	//PRODUCT
	c.productRepo = repository.NewProductRepository(c.db)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handlers.NewProductHandler(c.productService)

	// REVIEW PRODUCT
	c.reviewProductRepo = repository.NewReviewProductRepository(c.db)
	c.reviewProductSevice = service.NewReviewProductService(c.reviewProductRepo)
	c.reviewProductHandler = handlers.NewReviewProductHandler(c.reviewProductSevice)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}
func (c *Container) ForgotPwdHandle() *handlers.ForgotPwdHandler {
	return c.fpHandler
}
func (c *Container) LandingHandler() *handlers.LandingHandler {
	return c.landingHandler
}

func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}
func (c *Container) TransactionHandler() *handlers.TransactionHandler {
	return c.transactionHandler
}
func (c *Container) ReviewProductHandler() *handlers.ReviewProductHandler {
	return c.reviewProductHandler
}
