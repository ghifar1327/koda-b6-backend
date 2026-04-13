package repository

import (
	"backend/internal/dto"
	"backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type ProductRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewProductRepository(db *pgxpool.Pool, rdb *redis.Client) *ProductRepository {
	return &ProductRepository{
		db:  db,
		rdb: rdb,
	}
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	key := "product"
	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result []models.Product
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}

	query := `
		SELECT
		  p.id,
		  MIN(i.url) AS image,
		  p.name,
		  p.description,
		  p.price,
		  ARRAY_AGG(DISTINCT c.name) AS categories,
		  ARRAY_AGG(DISTINCT s.name) AS sizes,
		  ARRAY_AGG(DISTINCT v.name) AS variants,
		  COALESCE(AVG(rp.rating), 0) AS rating,
		  p.stock,
		  p.created_at,
		  p.updated_at
		FROM products p
		LEFT JOIN product_images pi 
			ON p.id = pi.product_id
		LEFT JOIN images i 
			ON pi.image_id = i.id
		LEFT JOIN products_categories pc 
			ON p.id = pc.product_id
		LEFT JOIN categories c 
			ON pc.category_id = c.id


		LEFT JOIN product_sizes ps
			ON p.id = ps.product_id
		LEFT JOIN sizes s
			ON s.id = ps.size_id


		LEFT JOIN product_variants pv
			ON p.id = pv.product_id
		LEFT JOIN variants v
			ON v.id = pv.variant_id


		LEFT JOIN transaction_details td 
			ON p.id = td.product_id
		LEFT JOIN review_product rp 
			ON td.id = rp.id_transaction_details
        GROUP BY p.id;`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(products)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	key := fmt.Sprintf("get-product-by-id:%d", id)
	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result models.Product
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}
	query := `
		SELECT
		  p.id,
		  MIN(i.url) AS image,
		  p.name,
		  p.description,
		  p.price,
		  ARRAY_AGG(DISTINCT c.name) AS categories,
		  ARRAY_AGG(DISTINCT s.name) AS sizes,
		  ARRAY_AGG(DISTINCT v.name) AS variants,
		  COALESCE(AVG(rp.rating), 0) AS rating,
		  p.stock,
		  p.created_at,
		  p.updated_at
		FROM products p
		LEFT JOIN product_images pi 
			ON p.id = pi.product_id
		LEFT JOIN images i 
			ON pi.image_id = i.id
		LEFT JOIN products_categories pc 
			ON p.id = pc.product_id
		LEFT JOIN categories c 
			ON pc.category_id = c.id


		LEFT JOIN product_sizes ps
			ON p.id = ps.product_id
		LEFT JOIN sizes s
			ON s.id = ps.size_id


		LEFT JOIN product_variants pv
			ON p.id = pv.product_id
		LEFT JOIN variants v
			ON v.id = pv.variant_id


		LEFT JOIN transaction_details td 
			ON p.id = td.product_id
		LEFT JOIN review_product rp 
			ON td.id = rp.id_transaction_details
		WHERE p.id = $1
		GROUP BY  p.id;`
	row, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	product, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(product)
    r.rdb.Set(ctx, key, data, 15*time.Minute)

	return &product, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p dto.CreateProductRequest) error {
	query := `INSERT INTO 
		products (
			name,
			description,
			price,
			stoct
			created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		time.Now(),
	)

	return err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id int, p models.Product) error {
	query := `
	    UPDATE products
	    SET 
	    	name=$1, 
	    	description=$2, 
	    	price=$3, 
	    	stock=$4, 
	    	cearted_at=$5 
	    WHERE id=$6`
	_, err := r.db.Exec(ctx, query,
		p.Name,
		p.Description,
		p.Price,
		p.Stock,
		time.Now(),
		id)

	r.rdb.Del(ctx, fmt.Sprintf("get-product-by-id:%d", id))
	return err
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM Products WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	r.rdb.Del(ctx, fmt.Sprintf("get-product-by-id:%d", id))
	return err
}

// =============================================================================================== GET SIZE AND VARIANT PRODUCT

func (r *ProductRepository) GetVariantsByIdProduct(ctx context.Context, id int) ([]models.Variant, error) {
	key := fmt.Sprintf("get-variant-product-by-id:%d", id)
	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result []models.Variant
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}
	query := `
		SELECT
			v.id,
			COALESCE(v.name, '') AS name,
	        COALESCE(v.add_price, 0) AS add_price
		FROM product_variants pv
		LEFT JOIN variants v ON pv.variant_id = v.id
		WHERE pv.product_id = $1`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	variants, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Variant])
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(variants)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}

	return variants, nil
}
func (r *ProductRepository) GetSizesByIdProduct(ctx context.Context, id int) ([]models.Size, error) {
	key := fmt.Sprintf("get-size-by-product-id:%d", id)
	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result []models.Size
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}
	
	query := `SELECT
		s.id,
	    COALESCE(s.name, '') AS name,
	    COALESCE(s.add_price, 0) AS add_price
	FROM product_sizes ps
	LEFT JOIN sizes s ON ps.size_id = s.id
	WHERE ps.product_id = $1`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sizes, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Size])
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(sizes)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}
	return sizes, nil
}
