package repository

import (
	"backend/internal/models"
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type LandingRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewLandingRepository(db *pgxpool.Pool, rdb *redis.Client) *LandingRepository {
	return &LandingRepository{
		db:  db,
		rdb: rdb,
	}
}

func (r *LandingRepository) GetAllReviewProductsLanding(ctx context.Context) ([]models.Reviews, error) {
	key := "landing/reviews"

	// cek redis
	cached , err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result []models.Reviews
		if err := json.Unmarshal([]byte(cached), &result); err == nil{
				return result, nil
		}
	}

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

	// simpan ke redis
	data , err := json.Marshal(reviews)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute * 15)
	}

	return reviews, nil
}
func (r *LandingRepository) GetReviewProductLandingByID(ctx context.Context, id int) (*models.Reviews, error) {
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
	
	key := "landing/recomended-product"

	cached , err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result []models.RecommendedProduct
		if err := json.Unmarshal([]byte(cached), &result); err == nil{
				return result, nil
		}
	}
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

	data , err := json.Marshal(recommended)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute * 15)
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

	row, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	rp, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.RecommendedProduct])
	if err != nil {
		return nil, err
	}

	return &rp, nil
}
