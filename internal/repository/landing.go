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
    	rp.rating
		FROM review_product rp
		JOIN transaction_details td ON rp.id_transaction_details = td.id
		JOIN products p ON td.product_id = p.id
		LEFT JOIN product_images pi ON p.id = pi.product_id
		LEFT JOIN images i ON pi.image_id = i.id;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reviews []models.Reviews
	for rows.Next() {
		var rp models.Reviews
		err := rows.Scan(
			&rp.Id,
			&rp.Name,
			&rp.Image,
			&rp.Description,
			&rp.Price,
			&rp.Rating,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, rp)
	}
	return reviews, nil
}

func (r *LandingRepository) GetReviwProductByID(ctx context.Context, id int) (*models.Reviews, error) {
	query := `
		SELECT 
    	p.id,
    	p.name,
    	i.url AS images,
    	p.description,
    	p.price,
    	rp.rating
		FROM review_product rp
		JOIN transaction_details td ON rp.id_transaction_details = td.id
		JOIN products p ON td.product_id = p.id
		LEFT JOIN product_images pi ON p.id = pi.product_id
		LEFT JOIN images i ON pi.image_id = i.id 
		WHERE id=$1`
	var rp models.Reviews
	err := r.db.QueryRow(ctx, query, id).Scan(
		&rp.Id,
		&rp.Name,
		&rp.Image,
		&rp.Description,
		&rp.Price,
		&rp.Rating,
	)
	if err != nil {
		return nil, err
	}
	return &rp, nil
}
// ======================================================================================================================== 	RECOMMENDED PRODUCT

func (r *LandingRepository) GetAllRecommendedProducts(ctx context.Context) ([]models.RecommendedProduct, error) {
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
	var recommended []models.RecommendedProduct
	for rows.Next() {
		var rp models.RecommendedProduct
		err := rows.Scan(
			&rp.Id,
			&rp.Name,
			&rp.Image,
			&rp.Description,
			&rp.Price,
			&rp.TotalReview,
			&rp.AvgRating,
		)
		if err != nil {
			return nil, err
		}
		recommended = append(recommended, rp)
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
    	rp.rating
		FROM review_product rp
		JOIN transaction_details td ON rp.id_transaction_details = td.id
		JOIN products p ON td.product_id = p.id
		LEFT JOIN product_images pi ON p.id = pi.product_id
		LEFT JOIN images i ON pi.image_id = i.id 
		WHERE id=$1`
	var rp models.RecommendedProduct
	err := r.db.QueryRow(ctx, query, id).Scan(
		&rp.Id,
		&rp.Name,
		&rp.Image,
		&rp.Description,
		&rp.Price,
		&rp.TotalReview,
		&rp.AvgRating,
	)
	if err != nil {
		return nil, err
	}
	return &rp, nil
}

