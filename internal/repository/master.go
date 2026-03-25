package repository

import (
	"backend/internal/dto"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type MasterRepository struct {
	db *pgxpool.Pool
	rdb *redis.Client
}

func NewMasterRepositoy(db *pgxpool.Pool, rdb *redis.Client) *MasterRepository {
	return &MasterRepository{db: db, rdb: rdb}
}

func (r *MasterRepository) Create(ctx context.Context, table string, req dto.CreateMasterRequest) error {
	query := fmt.Sprintf("INSERT INTO %s (name , add_price) VALUES ($1 ,$2) ", table)

	_, err := r.db.Exec(ctx, query, req.Name, req.AddPrice)

	return err
}

func (r *MasterRepository) GetAll(ctx context.Context, table string) ([]dto.Master, error) {
	key := table

	cached , err := r.rdb.Get(ctx,key).Result()
	if err == nil {
		var result []dto.Master
		if err := json.Unmarshal([]byte(cached), &result); err == nil{
			return result, nil
		}
	}

	query := fmt.Sprintf("SELECT id, name, add_price FROM %s", table)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results , err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Master])
	if err != nil {
		return nil, err
	}

	data , err := json.Marshal(results)
	if err == nil{
		r.rdb.Set(ctx, key, data, time.Minute * 15)
	}
	return results, nil
}

func (r *MasterRepository) GetById(ctx context.Context, table string, id int) (dto.Master, error) {
	query := fmt.Sprintf(`
		SELECT id, name, add_price 
		FROM %s 
		WHERE id=$1
	`, table)

	row, err := r.db.Query(ctx, query, id)
	if err != nil {
		return dto.Master{}, err
	}
	defer row.Close()
	data, err := pgx.CollectOneRow(row, pgx.RowToStructByName[dto.Master])
	if err != nil {
		return dto.Master{}, err
	}

	return data, nil
}

func (r *MasterRepository) Update(ctx context.Context, table string, id int, req dto.UpdateMasterRequest) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 , add_price = $2 WHERE id = $3 ", table)

	_, err := r.db.Exec(ctx, query, req.Name, req.AddPrice, id)

	if err != nil {
		return err
	}

	// hapus key dari redis
	r.rdb.Del(ctx, table)

	return nil
}

func (r *MasterRepository) Delete(ctx context.Context, table string, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table)

	_, err := r.db.Exec(ctx, query, id)
	
	if err != nil {
		return err
	}

	// hapus key dari redis
	r.rdb.Del(ctx, table)

	return nil
}
