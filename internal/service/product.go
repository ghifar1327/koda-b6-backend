package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error){
	return s.repo.GetAllProducts(ctx)
}

func (s *ProductService) GetAllProductByID(ctx context.Context , id int) (*models.Product, error){
	return s.repo.GetProductByID(ctx, id)
}