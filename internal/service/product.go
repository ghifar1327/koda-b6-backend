package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"errors"
	"strings"
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

func (s *ProductService) CreateProduct(ctx context.Context, req dto.CreateProductRequest) error {
	return s.repo.CreateProduct(ctx,req)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, req models.Product) error {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return err
	}

	if strings.TrimSpace(req.Name) != "" {
		product.Name = req.Name
	}

	if strings.TrimSpace(req.Description) != "" {
		product.Description = req.Description
	}

	if req.Stoct <= 0 {
		return errors.New("invalid Price")
	}
	product.Price = req.Price

	if req.Stoct <= 0 {
		return errors.New("invalid stock")
	}
	product.Stoct = req.Stoct

	return s.repo.UpdateProduct(ctx, id, *product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	return s.repo.DeleteProduct(ctx, id)
}
