package repository

import (
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
			id,
			name,
			description,
			price,
			stock,
			cteated_at,
			updated_atr
		FROM products`
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
			id,
			name,
			description,
			price,
			stock,
			cteated_at,
			updated_atr
		FROM products WHERE id=$1`
	var product models.Product
	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UploadedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
