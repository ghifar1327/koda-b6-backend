package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type ReviewProductRepository struct {
	db *pgx.Conn
}

func NewReviewProductRepository(db *pgx.Conn) *ReviewProductRepository {
	return &ReviewProductRepository{
		db: db,
	}
}

func (r *ReviewProductRepository) GetAllReviewProducts(ctx context.Context) ([]models.ReviewProduct, error) {
	query := `
		SELECT 
		id, 
		user_id,
		id_transaction_details,
		rating,
		message FROM review_product`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ReviewProducts []models.ReviewProduct
	for rows.Next() {
		var t models.ReviewProduct
		err := rows.Scan(
			t.Id,
			t.UserId,
			t.IdTransactionDetail,
			t.Rating,
			t.Message,
		)
		if err != nil {
			return nil, err
		}
		ReviewProducts = append(ReviewProducts, t)
	}
	return ReviewProducts, nil
}

// ==================================================================================================================================================== Get ReviewProduct By ID
func (r *ReviewProductRepository) GetReviewProductByID(ctx context.Context, id int) (*models.ReviewProduct, error) {
	query := `
		SELECT 
		id, 
		user_id,
		id_transaction_details,
		rating
		message
		FROM review_product WHERE id=$1`

	var ReviewProduct models.ReviewProduct

	err := r.db.QueryRow(ctx, query, id).Scan(
		&ReviewProduct.Id,
		&ReviewProduct.UserId,
		&ReviewProduct.IdTransactionDetail,
		&ReviewProduct.Rating,
		&ReviewProduct.Message,
	)

	if err != nil {
		return nil, err
	}

	return &ReviewProduct, nil
}

// ====================================================================================================================================================  Create ReviewProduct

func (r *ReviewProductRepository) CreateReviewProduct(ctx context.Context, t models.ReviewProduct) error {
	query := `INSERT INTO review_product (
		id, 
		user_id,
		id_transaction_details,
		rating,
		message) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query,
		t.Id,
		t.UserId,
		t.IdTransactionDetail,
		t.Rating,
		t.Message)

	return err
}

// ==================================================================================================================================================== Update ReviewProduct

func (r *ReviewProductRepository) UpdateReviewProduct(ctx context.Context, id int, t models.ReviewProduct) error {
	query := `
	    UPDATE review_product SET rating=$1, message=$2  WHERE id=$3`
	_, err := r.db.Exec(ctx, query, t, t.Rating, t.Message, id)
	return err
}

// ======================================================================================================== DELETE ReviewProduct
func (r *ReviewProductRepository) DeleteReviewProduct(ctx context.Context, id int) error {
	query := `DELETE FROM review_product WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
