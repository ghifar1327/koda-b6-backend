package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"context"

	"github.com/google/uuid"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddToCart(ctx context.Context, req dto.ADDCartRequest) ([]dto.ADDCartRequest, error) {
	existing, err := s.repo.FindExisting(ctx, req)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		newQty := existing.Quantity + req.Quantity
		return s.repo.UpdateQuantity(ctx, existing.Id, newQty)
	}

	return s.repo.AddCart(ctx, req)
}

func (s *CartService) GetCartByUserId(ctx context.Context, userID uuid.UUID) ([]models.Cart, error) {
	return s.repo.GetCartByUserId(ctx, userID)
}

func (s *CartService) DeleteCart(ctx context.Context, id int) ([]models.Cart ,error) {
	return s.repo.Delete(ctx, id)
}
