package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"time"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetAllProducts(ctx)
}

func (s *ProductService) GetAllProductByID(ctx context.Context, id int) (*models.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) CreateProduct(ctx context.Context, req dto.CreteProductRequest) error {
	newProduct := dto.CreteProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stoct:       req.Stoct,
		CreatedAt:   time.Now(),
	}
	return s.repo.CreateProduct(ctx, newProduct)
}
