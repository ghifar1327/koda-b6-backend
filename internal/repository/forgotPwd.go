package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type DataRepository struct {
	db *pgx.Conn
}

func NewDataRepository(db *pgx.Conn) *DataRepository {
	return &DataRepository{
		db: db,
	}
}

func (r *DataRepository) CreateData(ctx context.Context, f models.ForgotPassword) error {
	query := `INSERT INTO forgot_pwd (email, code) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, f.Email, f.Code)

	return err
}

func (r *DataRepository) GetDataByEmail(ctx context.Context, email string) (*models.ForgotPassword, error) {
	query := `
		SELECT 
		email,
		code
		FROM forgot_pwd WHERE email=$1`

	var fp models.ForgotPassword

	err := r.db.QueryRow(ctx, query, email).Scan(
		fp.Email,
		fp.Code,
	)

	if err != nil {
		return nil, err
	}

	return &fp, nil
}

func (r *DataRepository) DeleteDataByCode(ctx context.Context, code int) error {
	query := `DELETE FROM forgot_pwd WHERE code=$1`
	_, err := r.db.Exec(ctx, query, code)
	return err
}
