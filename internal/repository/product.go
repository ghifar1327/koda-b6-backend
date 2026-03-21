package repository

import (
	"backend/internal/dto"
	"backend/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	query := `
		SELECT
		  p.id,
		  MIN(i.url) AS image,
		  p.name,
		  p.description,
		  p.price,
		  ARRAY_AGG(DISTINCT c.name) AS categories,
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
		LEFT JOIN transaction_details td 
			ON p.id = td.product_id
		LEFT JOIN review_product rp 
			ON td.id = rp.id_transaction_details
        GROUP BY p.id;`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT
		  p.id,
		  MIN(i.url) AS image,
		  p.name,
		  p.description,
		  p.price,
		  ARRAY_AGG(DISTINCT c.name) AS categories,
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
		LEFT JOIN transaction_details td 
			ON p.id = td.product_id
		LEFT JOIN review_product rp 
			ON td.id = rp.id_transaction_details
		WHERE p.id = $1
		GROUP BY  p.id;`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p dto.CreateProductRequest) error {
	query := `INSERT INTO 
		products (
			name,
			description,
			price,
			stoct,
			crated_at) VALUES ($1, $2, $3, $4, $5)`
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
	return err
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM Products WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
