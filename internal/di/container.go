package di

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/service"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	db  *pgxpool.Pool
	rdb *redis.Client

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

	// CART
	cartREPO    *repository.CartRepository
	cartService *service.CartService
	cartHandler *handlers.CartHandler

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
	masterRepo    *repository.MasterRepository
	masterService *service.MasterService
	masterHandler *handlers.MasterHandler
}

func NewContainer(db *pgxpool.Pool) *Container {
	godotenv.Load()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	container := &Container{
		db:  db,
		rdb: redisClient,
	}

	container.initDependencies()

	return container
}

func (c *Container) initDependencies() {

	// USER
	c.userRepo = repository.NewUserrepository(c.db, c.rdb)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)

	//AUTH
	c.fpRepo = repository.NewAuthRepository(c.db, c.rdb)
	c.fpService = service.NewAuthService(c.userRepo, c.fpRepo)
	c.authHandler = handlers.NewAuthHandler(c.fpService, c.userService)

	//LANDING
	c.landingRepo = repository.NewLandingRepository(c.db, c.rdb)
	c.landingService = service.NewLandingService(c.landingRepo)
	c.landingHandler = handlers.NewLandingHandler(c.landingService)

	//CARD
	c.cartREPO = repository.NewCartRepository(c.db, c.rdb)
	c.cartService = service.NewCartService(c.cartREPO)
	c.cartHandler = handlers.NewCartHandler(c.cartService)

	// TRANSACTION
	c.transactionRepo = repository.NewTransactionRepository(c.db, c.rdb)
	c.transactionService = service.NewTransactionService(c.transactionRepo)
	c.transactionHandler = handlers.NewTransactionHandler(c.transactionService)

	//PRODUCT
	c.productRepo = repository.NewProductRepository(c.db, c.rdb)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handlers.NewProductHandler(c.productService)

	// REVIEW PRODUCT
	c.reviewProductRepo = repository.NewReviewProductRepository(c.db, c.rdb)
	c.reviewProductSevice = service.NewReviewProductService(c.reviewProductRepo)
	c.reviewProductHandler = handlers.NewReviewProductHandler(c.reviewProductSevice)

	//MASTER
	c.masterRepo = repository.NewMasterRepositoy(c.db, c.rdb)
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
func (c *Container) CartHandler() *handlers.CartHandler {
	return c.cartHandler
}
func (c *Container) TransactionHandler() *handlers.TransactionHandler {
	return c.transactionHandler
}
func (c *Container) ReviewProductHandler() *handlers.ReviewProductHandler {
	return c.reviewProductHandler
}

func (c *Container) MasterHandler() *handlers.MasterHandler {
	return c.masterHandler
}
