package service

import (
	"context"
	"backend/internal/dto"
	"backend/internal/repository"
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

func (s *CartService) GetCart(ctx context.Context, userID string) ([]dto.ADDCartRequest, error) {
	return s.repo.GetCartByUser(ctx, userID)
}

func (s *CartService) DeleteCart(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}