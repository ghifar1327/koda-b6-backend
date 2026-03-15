package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"strings"
	"github.com/google/uuid"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) GetAllTransactions(ctx context.Context) ([]models.Transaction, error) {
	return s.repo.GetAllTransaction(ctx)
}

func (s *TransactionService) GetAllTransactionByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	return s.repo.GetTransactionByID(ctx, id)
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) error {
	return s.repo.CreateTransaction(ctx, req)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id uuid.UUID, req models.Transaction) error {
	Transaction, err := s.repo.GetTransactionByID(ctx, id)
	if err != nil {
		return err
	}

	if strings.TrimSpace(req.Status) != "" {
		Transaction.Status = req.Status
	}
	newTransaction := Transaction.Status

	return s.repo.UpdateTransaction(ctx, id, newTransaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteTransaction(ctx, id)
}
