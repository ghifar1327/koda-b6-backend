package repository

import (
	"backend/internal/dto"
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReviewProductRepository struct {
	db *pgxpool.Pool
}

func NewReviewProductRepository(db *pgxpool.Pool) *ReviewProductRepository {
	return &ReviewProductRepository{
		db: db,
	}
}

// ======================================================================================================== GET ALL REVIEW PRODUCTS

func (r *ReviewProductRepository) GetAllReviewProducts(ctx context.Context) ([]models.ReviewProduct, error) {
	query := `
		SELECT 
			id, 
			user_id,
			id_transaction_details,
			rating,
			message 
		FROM review_product
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	reviewProducts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ReviewProduct])
	if err != nil {
		return nil, err
	}

	return reviewProducts, nil
}

// ======================================================================================================== GET REVIEW PRODUCT BY ID

func (r *ReviewProductRepository) GetReviewProductByID(ctx context.Context, id int) (*models.ReviewProduct, error) {
	query := `
		SELECT 
			id, 
			user_id,
			id_transaction_details,
			rating,
			message
		FROM review_product
		WHERE id=$1
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	reviewProduct, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ReviewProduct])
	if err != nil {
		return nil, err
	}

	return &reviewProduct, nil
}

// ======================================================================================================== CREATE REVIEW PRODUCT

func (r *ReviewProductRepository) CreateReviewProduct(ctx context.Context, rp dto.CreateReviewProductRequest) error {
	query := `
		INSERT INTO review_product (
			user_id,
			id_transaction_details,
			rating,
			message
		) 
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query,
		rp.UserId,
		rp.IdTransactionDetail,
		rp.Rating,
		rp.Message,
	)

	return err
}

// ======================================================================================================== UPDATE REVIEW PRODUCT

func (r *ReviewProductRepository) UpdateReviewProduct(ctx context.Context, id int, t models.ReviewProduct) error {
	query := `
		UPDATE review_product 
		SET rating=$1, message=$2
		WHERE id=$3
	`

	_, err := r.db.Exec(ctx, query,
		t.Rating,
		t.Message,
		id,
	)

	return err
}

// ======================================================================================================== DELETE REVIEW PRODUCT

func (r *ReviewProductRepository) DeleteReviewProduct(ctx context.Context, id int) error {
	query := `
		DELETE FROM review_product 
		WHERE id=$1
	`

	_, err := r.db.Exec(ctx, query, id)
	return err
}
