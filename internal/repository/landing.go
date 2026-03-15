package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type LandingRepository struct {
	db *pgx.Conn
}

func NewLandingRepository(db *pgx.Conn) *LandingRepository {
	return &LandingRepository{
		db: db,
	}
}

func (r *LandingRepository) GetAllReviewProducs(ctx context.Context) ([]models.Reviews, error) {
	query := `
		SELECT 
            p.id,
            p.name,
            i.url AS images,
            p.description,
            p.price,
            COUNT(rp.id) AS total_review
        FROM review_product rp
        JOIN transaction_details td ON rp.id_transaction_details = td.id
        JOIN products p ON td.product_id = p.id
        LEFT JOIN product_images pi ON p.id = pi.product_id
        LEFT JOIN images i ON pi.image_id = i.id
        GROUP BY p.id,p.name,i.url,p.description,p.price
        ORDER BY total_review DESC;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Reviews])
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
func (r *LandingRepository) GetReviwProductByID(ctx context.Context, id int) (*models.Reviews, error) {
	query := `
		SELECT 
        p.id,
        p.name,
        i.url AS image,
        p.description,
        p.price,
        COUNT(rp.id) AS total_review
        FROM review_product rp
        JOIN transaction_details td ON rp.id_transaction_details = td.id
        JOIN products p ON td.product_id = p.id
        LEFT JOIN product_images pi ON p.id = pi.product_id
        LEFT JOIN images i ON pi.image_id = i.id
		WHERE p.id=$1
        GROUP BY p.id,p.name,i.url,p.description,p.price;`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	review, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Reviews])
	if err != nil {
		return nil, err
	}

	return &review, nil
}

// ======================================================================================================================================================== RECOMMENDED PRODUCT

func (r *LandingRepository) GetRecommendedProducts(ctx context.Context) ([]models.RecommendedProduct, error) {
	query := `
		SELECT 
            p.id,
            p.name,
            i.url AS images,
            p.description,
            p.price,
            COUNT(rp.id) AS total_review,
            AVG(rp.rating) AS avg_rating
  	    FROM review_product rp
  	    JOIN transaction_details td ON rp.id_transaction_details = td.id
  	    JOIN products p ON td.product_id = p.id
  	    LEFT JOIN product_images pi ON p.id = pi.product_id
  	    LEFT JOIN images i ON pi.image_id = i.id
  	    GROUP BY p.id, p.name, i.url, p.description, p.price
  	    ORDER BY avg_rating DESC
		LIMIT 4;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	recommended, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.RecommendedProduct])
	if err != nil {
		return nil, err
	}

	return recommended, nil
}

func (r *LandingRepository) GetRecommendedProductByID(ctx context.Context, id int) (*models.RecommendedProduct, error) {

	query := `
	SELECT 
        p.id,
        p.name,
        i.url AS images,
        p.description,
        p.price,
        COUNT(DISTINCT rp.id) AS total_review,
        AVG(rp.rating) AS avg_rating
	FROM review_product rp
	JOIN transaction_details td ON rp.id_transaction_details = td.id
	JOIN products p ON td.product_id = p.id
	LEFT JOIN product_images pi ON p.id = pi.product_id
	LEFT JOIN images i ON pi.image_id = i.id
	WHERE p.id=$1
	GROUP BY p.id, p.name, i.url, p.description, p.price
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rp, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.RecommendedProduct])
	if err != nil {
		return nil, err
	}

	return &rp, nil
}
