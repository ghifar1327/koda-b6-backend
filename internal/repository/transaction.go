package repository

import (
	"backend/internal/dto"
	"backend/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type TransactionRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewTransactionRepository(db *pgxpool.Pool, redis *redis.Client) *TransactionRepository {
	return &TransactionRepository{
		db:  db,
		rdb: redis,
	}
}

func (r *TransactionRepository) GetAllTransaction(ctx context.Context) ([]models.Transaction, error) {
	key := "get-all-transaction"
	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result []models.Transaction
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}

	query := `SELECT
	t.id,
	u.full_name,
	t.address,
	u.phone,
	m.name AS shipping,
	t.payment_method,
	t.status,
	t.created_at,
	t.updeted_at,

	p.name AS product_name,
	img.url AS product_image,
	s.name AS size,
	v.name AS variant,
	td.quantity,

	(p.price + COALESCE(s.add_price,0) + COALESCE(v.add_price,0)) * td.quantity AS subtotal

	FROM transactions t

	JOIN users u ON u.id = t.user_id
	JOIN methods m ON m.id = t.id_method

	JOIN transaction_details td ON td.transaction_id = t.id
	JOIN products p ON p.id = td.product_id

	LEFT JOIN sizes s ON s.id = td.size_id
	LEFT JOIN variants v ON v.id = td.variant_id

	LEFT JOIN product_images pi ON pi.product_id = p.id
	LEFT JOIN images img ON img.id = pi.image_id

	ORDER BY t.created_at DESC;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionRow])
	if err != nil {
		return nil, err
	}

	transactionMap := map[uuid.UUID]*models.Transaction{}

	for _, r := range result {

		trx, exists := transactionMap[r.Id]

		if !exists {
			trx = &models.Transaction{
				Id:            r.Id,
				FullName:      r.FullName,
				Address:       r.Address,
				Phone:         r.Phone,
				Shipping:      r.Shipping,
				PaymentMethod: r.PaymentMethod,
				Status:        r.Status,
				CreatedAt:     r.CreatedAt,
				Items:         []models.ItemDetail{},
			}

			transactionMap[r.Id] = trx
		}

		item := models.ItemDetail{
			TransactionId: r.Id,
			ProductImage:  r.ProductImage,
			ProductName:   r.ProductName,
			Size:          r.Size,
			Variant:       r.Variant,
			Quantity:      r.Quantity,
			SubTotal:      r.Subtotal,
		}

		trx.Items = append(trx.Items, item)
		trx.TotalTransaction += r.Subtotal
	}

	var transactions []models.Transaction

	for _, v := range transactionMap {
		transactions = append(transactions, *v)
	}

	data, err := json.Marshal(transactions)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}

	return transactions, nil
}

// ====================================================================================================================================== Get Transaction By ID
func (r *TransactionRepository) GetTransactionByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {

	key := fmt.Sprintf("get-transaction-by-id:%s", id.String())
	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result models.Transaction
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}

	query := `
	SELECT
	t.id,
	u.full_name,
	t.address,
	u.phone,
	m.name AS shipping,
	t.payment_method,
	t.status,
	t.created_at,
	t.updeted_at,


	p.name AS product_name,
	img.url AS product_image,
	s.name AS size,
	v.name AS variant,
	td.quantity,

	(p.price + COALESCE(s.add_price,0) + COALESCE(v.add_price,0)) * td.quantity AS subtotal

	FROM transactions t

	JOIN users u ON u.id = t.user_id
	JOIN methods m ON m.id = t.id_method

	JOIN transaction_details td ON td.transaction_id = t.id
	JOIN products p ON p.id = td.product_id

	LEFT JOIN sizes s ON s.id = td.size_id
	LEFT JOIN variants v ON v.id = td.variant_id

	LEFT JOIN product_images pi ON pi.product_id = p.id
	LEFT JOIN images img ON img.id = pi.image_id

	WHERE t.id = $1;
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionRow])
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, err
	}

	trx := models.Transaction{
		Id:            result[0].Id,
		FullName:      result[0].FullName,
		Address:       result[0].Address,
		Phone:         result[0].Phone,
		Shipping:      result[0].Shipping,
		PaymentMethod: result[0].PaymentMethod,
		Status:        result[0].Status,
		CreatedAt:     result[0].CreatedAt,
		UpdatedAt:     result[0].UpdatedAt,
	}

	total := 0

	for _, r := range result {

		item := models.ItemDetail{
			TransactionId: r.Id,
			ProductName:   r.ProductName,
			ProductImage:  r.ProductImage,
			Size:          r.Size,
			Variant:       r.Variant,
			Quantity:      r.Quantity,
			SubTotal:      r.Subtotal,
		}

		total += r.Subtotal

		trx.Items = append(trx.Items, item)
	}

	trx.TotalTransaction = total

	data, err := json.Marshal(trx)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}

	return &trx, nil
}

// ====================================================================================================================================================  Create Transaction

func (r *TransactionRepository) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	IdTransaction := uuid.New()

	query := `INSERT INTO transactions (
		id, 
		user_id,
		address,
		status,
		id_method,
		payment_method,
		id_voucher,
		created_at,
		updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = tx.Exec(ctx, query,
		IdTransaction,
		req.UserId,
		req.Address,
		"pending",
		req.IdMethod,
		req.PaymentMethod,
		req.IdVoucher,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	for _, item := range req.Items {
		queryDetail := `INSERT INTO transaction_details(
		transaction_id,
		product_id,
		size_id,
		variant_id,
		quantity
		) VALUES ($1, $2, $3, $4, $5)`

		_, err = tx.Exec(ctx, queryDetail,
			IdTransaction,
			item.ProductId,
			item.SizeId,
			item.VariantId,
			item.Quantity,
		)
		if err != nil {
			return err
		}

		updateStock := `
		UPDATE products
		SET stock = stock - $1
		WHERE id = $2 AND stock >= $1
		`
		result, err := tx.Exec(ctx, updateStock,
			item.Quantity,
			item.ProductId,
		)

		if err != nil {
			return err
		}
		if result.RowsAffected() == 0 {
			return errors.New("Stock not enough")
		}
	}

	return tx.Commit(ctx)
}

// ==================================================================================================================================================== Update Transaction

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, id uuid.UUID, status string) error {
	query := `
	    UPDATE Transactions SET status=$1 ,updated_at = $2, WHERE id = $3`
	_, err := r.db.Exec(ctx, query, status, time.Now(), id)
	r.rdb.Del(ctx, fmt.Sprintf("get-transaction-by-id:%s", id.String()))
	return err
}

// ======================================================================================================== DELETE Transaction
func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM Transactions WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	r.rdb.Del(ctx, fmt.Sprintf("get-transaction-by-id:%s", id.String()))
	return err
}
