package repository

import (
	"backend/internal/dto"
	"backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	cacheKey := fmt.Sprintf(
		"cart:item:%s:%d:%d:%d",
		req.UserID.String(),
		req.ProductID,
		req.SizeID,
		req.Variant,
	)

	val, err := r.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var cached dto.ADDCartRequest
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			return &cached, nil
		}
	}

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

	result := data[0]

	jsonData, _ := json.Marshal(result)
	r.rdb.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return &result, nil
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
	r.rdb.Del(ctx, "cart:"+req.UserID.String())

	return data, nil
}

func (r *CartRepository) UpdateQuantity(ctx context.Context, id int, qty int) ([]dto.ADDCartRequest, error) {
	var userID uuid.UUID

	// ambil user_id
	err := r.db.QueryRow(ctx,
		`SELECT user_id FROM cart WHERE id=$1`,
		id,
	).Scan(&userID)
	if err != nil {
		return nil, err
	}

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

	r.rdb.Del(ctx, "cart:"+userID.String())

	return data, nil
}

func (r *CartRepository) GetCartByUserId(ctx context.Context, userID uuid.UUID) ([]models.Cart, error) {
	cacheKey := "cart:" + userID.String()

	val, err := r.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var cached []models.Cart
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			return cached, nil
		}
	}

	query := `
	SELECT 
	 c.id,
     i.url AS product_image,
     p.name AS product_name,
     s.name AS size,
     v.name AS variant,
     c.quantity,
     ((p.price + COALESCE(s.add_price, 0) + COALESCE(v.add_price, 0)) * c.quantity) AS subtotal
    FROM cart c
    JOIN products p ON c.product_id = p.id
    LEFT JOIN sizes s ON c.size_id = s.id
    LEFT JOIN variants v ON c.variant_id = v.id
    LEFT JOIN product_images pi ON p.id = pi.product_id
    LEFT JOIN images i ON pi.image_id = i.id
    WHERE c.user_id = $1;
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Cart])
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(data)
	r.rdb.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return data, nil
}

func (r *CartRepository) Delete(ctx context.Context, id int) ([]models.Cart, error) {
	var userID uuid.UUID

	err := r.db.QueryRow(ctx,
		`SELECT user_id FROM cart WHERE id = $1`,
		id,
	).Scan(&userID)

	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx,
		`DELETE FROM cart WHERE id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}

	r.rdb.Del(ctx, "cart:"+userID.String())
	return r.GetCartByUserId(ctx, userID)
}
