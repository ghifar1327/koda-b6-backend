package repository

import (
	"backend/internal/dto"
	"backend/internal/models"
	"context"

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
			p.name,
			p.description,
			p.price,
			ARRAY_AGG(c.name) AS categories,
			p.stock,
			p.created_at,
			p.updated_at
		FROM products p
		JOIN products_categories pc
			ON p.id = pc.product_id
		JOIN categories c
			ON pc.category_id = c.id
		GROUP BY p.id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Categories,
			&p.Stoct,
			&p.CreatedAt,
			&p.UploadedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.description,
			p.price,
			ARRAY_AGG(c.name) AS categories,
			p.stock,
			p.created_at,
			p.updated_at
		FROM products p
		JOIN products_categories pc
			ON p.id = pc.product_id
		JOIN categories c
			ON pc.category_id = c.id
		GROUP BY p.id WHERE id=$1`
	var product models.Product
	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Categories,
		&product.Stoct,
		&product.CreatedAt,
		&product.UploadedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p dto.CreteProductRequest) error {
	query := `INSERT INTO 
		products (
			name,
			description,
			price,
			stoct,
			crated_at) VALUES (1$, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query,
		p.Name,
		p.Description,
		p.Price,
		p.Stoct,
		p.CreatedAt,
	)

	return err
}
