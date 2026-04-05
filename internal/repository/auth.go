package repository

import (
	"backend/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewAuthRepository(db *pgxpool.Pool, rdb *redis.Client) *AuthRepository {
	return &AuthRepository{
		db:  db,
		rdb: rdb,
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

func (r *AuthRepository) UpdatePassword(ctx context.Context, id uuid.UUID, newPassword string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := r.db.Exec(ctx, query, newPassword, id)
	return err
}

func (r *AuthRepository) DeleteForgotPWDByCode(ctx context.Context, code int) error {
	query := `DELETE FROM forgot_password WHERE code=$1`
	_, err := r.db.Exec(ctx, query, code)
	return err
}


func (r *UserRepository) UpdateUserPicture(ctx context.Context, id uuid.UUID, picture string) (models.User, error) {
	query := `
	    UPDATE users 
	    SET 
	    	picture = $1,
	    	updated_at = $2
	    WHERE id = $3
	`

	_, err := r.db.Exec(ctx, query,
		picture,
		time.Now(),
		id,
	)
	if err != nil {
		return models.User{}, err
	}

	// clear cache
	r.rdb.Del(ctx, fmt.Sprintf("user:id:%s", id.String()))

	user, err := r.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}