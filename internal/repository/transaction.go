package repository

import (
	"backend/internal/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	db *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) GetAllTransaction(ctx context.Context) ([]models.Transaction, error) {
	query := `
		SELECT 
			id, 
			user_id,
			status,
			id_methode,
			payment_method,
			id_voucher,
			created_at
			FROM Transactions`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(
			t.Id,
			t.UserId,
			t.Status,
			t.IdMethod,
			t.PaymentMethode,
			t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		Transactions = append(Transactions, t)
	}
	return Transactions, nil
}

// ==================================================================================================================================================== Get Transaction By ID
func (r *TransactionRepository) GetTransactionByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	query := `
		SELECT 
		id, 
		user_id,
		status,
		id_methode,
		payment_method,
		id_voucher,
		created_at
		FROM Transactions WHERE id=$1`

	var Transaction models.Transaction

	err := r.db.QueryRow(ctx, query, id).Scan(
		&Transaction.Id,
		&Transaction.UserId,
		&Transaction.Status,
		&Transaction.IdMethod,
		&Transaction.PaymentMethode,
		&Transaction.IdVoucher,
		&Transaction.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &Transaction, nil
}

// ====================================================================================================================================================  Create Transaction

func (r *TransactionRepository) CreateTransaction(ctx context.Context, t models.Transaction) error {
	query := `INSERT INTO Transactions (
		id, 
		user_id,
		status,
		id_methode,
		payment_method,
		id_voucher,
		created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query,
		t.Id,
		t.UserId,
		t.Status,
		t.IdMethod,
		t.PaymentMethode,
		t.IdVoucher,
		t.CreatedAt)

	return err
}

// ==================================================================================================================================================== Update Transaction

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, id uuid.UUID, t models.Transaction) error {
	query := `
	    UPDATE Transactions SET status=$1, WHERE id=$2`
	_, err := r.db.Exec(ctx, query, t.Status, id)
	return err
}

// ======================================================================================================== DELETE Transaction
func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM Transactions WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// func (s *TransactionService) Register(ctx context.Context, req models.CreateTransactionRequest) error {
// 	if err := validateTransaction(req.Fullname, req.Email, req.Password); err != nil {
// 		return err
// 	}

// 	argon := argon2.DefaultConfig()
// 	encoded, err := argon.HashEncoded([]byte(req.Password))

// 	if err != nil {
// 		return err
// 	}

// 	newTransaction := models.Transaction{
// 		Fullname: req.Fullname,
// 		Email:    req.Email,
// 		Password: string(encoded),
// 	}
// 	return s.repo.Create(ctx, newTransaction)
// }
