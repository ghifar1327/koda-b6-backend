package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
)

type LandingService struct {
	repo *repository.LandingRepository
}

func NewLandingService(repo *repository.LandingRepository) *LandingService {
	return &LandingService{
		repo: repo,
	}
}

func (s *LandingService) GetAllReviewProductsLanding(ctx context.Context) ([]models.Reviews, error) {
	return s.repo.GetAllReviewProductsLanding(ctx)
}

func (s *LandingService) GetReviewProductLandingByID(ctx context.Context, id int) (*models.Reviews, error) {
	return s.repo.GetReviewProductLandingByID(ctx, id)
}


func (s *LandingService) GetRecommendedProducts(ctx context.Context) ([]models.RecommendedProduct, error) {
	return s.repo.GetRecommendedProducts(ctx)
}

func (s *LandingService) GetRecommendedProductByID(ctx context.Context, id int) (*models.RecommendedProduct, error) {
	return s.repo.GetRecommendedProductByID(ctx, id)
}