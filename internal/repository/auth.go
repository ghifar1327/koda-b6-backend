package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateForgotPWD(ctx context.Context, f models.ForgotPassword) error {
	query := `INSERT INTO forgot_password (email, code) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, f.Email, f.Code)

	return err
}

func (r *AuthRepository) GetForgotPWDByEmail(ctx context.Context, email string) (*models.ForgotPassword, error) {
	query := `
		SELECT 
		email,
		code
		FROM forgot_password WHERE email=$1`

	var fp models.ForgotPassword

	err := r.db.QueryRow(ctx, query, email).Scan(
		&fp.Email,
		&fp.Code,
	)

	if err != nil {
		return nil, err
	}

	return &fp, nil
}

func (r *AuthRepository) DeleteForgotPWDByCode(ctx context.Context, code int) error {
	query := `DELETE FROM forgot_password WHERE code=$1`
	_, err := r.db.Exec(ctx, query, code)
	return err
}
