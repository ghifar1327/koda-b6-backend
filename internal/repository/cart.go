package repository

import (
	"backend/internal/dto"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type CartRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewCartRepository(db *pgxpool.Pool, rdb *redis.Client) *CartRepository {
	return &CartRepository{db: db, rdb: rdb}
}

func (r *CartRepository) FindExisting(ctx context.Context, req dto.ADDCartRequest) (*dto.ADDCartRequest, error) {
	query := `
	SELECT id, user_id, product_id, size_id, variant_id, quantity
	FROM cart
	WHERE user_id=$1 AND product_id=$2 AND size_id=$3 AND variant_id=$4
	`

	rows, err := r.db.Query(ctx, query,
		req.UserID, req.ProductID, req.SizeID, req.Variant,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.ADDCartRequest])
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	return &data[0], nil
}

func (r *CartRepository) AddCart(ctx context.Context, req dto.ADDCartRequest) ([]dto.ADDCartRequest, error) {
	query := `
	INSERT INTO cart (user_id, product_id, size_id, variant_id, quantity)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id, user_id, product_id, size_id, variant_id, quantity
	`

	rows, err := r.db.Query(ctx, query,
		req.UserID, req.ProductID, req.SizeID, req.Variant, req.Quantity,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.ADDCartRequest])
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *CartRepository) UpdateQuantity(ctx context.Context, id int, qty int) ([]dto.ADDCartRequest, error) {
	query := `
	UPDATE cart
	SET quantity = $1
	WHERE id = $2
	RETURNING id, user_id, product_id, size_id, variant_id, quantity
	`

	rows, err := r.db.Query(ctx, query, qty, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.ADDCartRequest])
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *CartRepository) GetCartByUser(ctx context.Context, userID string) ([]dto.ADDCartRequest, error) {
	query := `
	SELECT id, user_id, product_id, size_id, variant_id, quantity
	FROM cart
	WHERE user_id=$1
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[dto.ADDCartRequest])
}

func (r *CartRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM cart WHERE id=$1`, id)
	return err
}
