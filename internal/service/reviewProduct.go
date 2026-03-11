package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"strings"
)

type ReviewProductService struct {
	repo *repository.ReviewProductRepository
}

func NewReviewProductService(repo *repository.ReviewProductRepository) *ReviewProductService {
	return &ReviewProductService{
		repo: repo,
	}
}

func (s *ReviewProductService) GetAllReviewProducts(ctx context.Context) ([]models.ReviewProduct, error) {
	return s.repo.GetAllReviewProducts(ctx)
}

func (s *ReviewProductService) GetAllReviewProductByID(ctx context.Context, id int) (*models.ReviewProduct, error) {
	return s.repo.GetReviewProductByID(ctx, id)
}

func (s *ReviewProductService) CreateReviewProduct(ctx context.Context, req models.ReviewProduct) error {
	newReviewProduct := models.ReviewProduct{
		Id:                  req.Id,
		UserId:              req.UserId,
		IdTransactionDetail: req.IdTransactionDetail,
		Rating:              req.Rating,
	}
	return s.repo.CreateReviewProduct(ctx, newReviewProduct)
}

func (s *ReviewProductService) UpdateReviewProduct(ctx context.Context, id int, req dto.UpdateReviewProductRequest) error {
	ReviewProduct, err := s.repo.GetReviewProductByID(ctx, id)
	if err != nil {
		return err
	}

	if strings.TrimSpace(req.Message) != "" {
		ReviewProduct.Message = req.Message
	}
	if req.Rating > 0 && req.Rating <= 10 {
		ReviewProduct.Rating = req.Rating
	}

	return s.repo.UpdateReviewProduct(ctx, id, *ReviewProduct)
}

func (s *ReviewProductService) DeleteReviewProduct(ctx context.Context, id int) error {
	return s.repo.DeleteReviewProduct(ctx, id)
}
