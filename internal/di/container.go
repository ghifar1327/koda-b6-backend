package di

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db *pgxpool.Pool

	// USER
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler

	//FORGOT PASSWORD
	fpRepo      *repository.AuthRepository
	fpService   *service.AuthService
	authHandler *handlers.AuthHandler

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

	// MASTER 
	masterRepo *repository.MasterRepository
	masterService *service.MasterService
	masterHandler *handlers.MasterHandler
}

func NewContainer(db *pgxpool.Pool) *Container {
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

	//AUTH
	c.fpRepo = repository.NewAuthRepository(c.db)
	c.fpService = service.NewAuthService(c.userRepo, c.fpRepo)
	c.authHandler = handlers.NewAuthHandler(c.fpService)

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

	//MASTER 
	c.masterRepo = repository.NewMasterRepositoy(c.db)
	c.masterService = service.NewMasterService(c.masterRepo)
	c.masterHandler = handlers.NewMasterHandler(c.masterService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}
func (c *Container) AuthHandler() *handlers.AuthHandler {
	return c.authHandler
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

func (c *Container) MasterHandler() *handlers.MasterHandler{
	return c.masterHandler	
}