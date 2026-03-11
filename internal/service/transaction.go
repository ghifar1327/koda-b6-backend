package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"strings"
	"time"

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

func (s *TransactionService) CreateTransaction(ctx context.Context, req models.Transaction) error {
	newTransaction := models.Transaction{
		Id:             uuid.New(),
		UserId:         req.UserId,
		Status:         req.Status,
		IdMethod:       req.IdMethod,
		PaymentMethode: req.PaymentMethode,
		IdVoucher:      req.IdVoucher,
		CreatedAt:      time.Now(),
	}
	return s.repo.CreateTransaction(ctx, newTransaction)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, id uuid.UUID, req models.Transaction) error {
	Transaction, err := s.repo.GetTransactionByID(ctx, id)
	if err != nil {
		return err
	}

	if strings.TrimSpace(req.Status) != "" {
		Transaction.Status = req.Status
	}

	return s.repo.UpdateTransaction(ctx, id, *Transaction)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteTransaction(ctx, id)
}
