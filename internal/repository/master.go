package repository

import (
	"backend/internal/dto"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MasterRepository struct {
	db *pgxpool.Pool
}

func NewMasterRepositoy(db *pgxpool.Pool) *MasterRepository {
	return &MasterRepository{db: db}
}

func (r *MasterRepository) Create(ctx context.Context, table string, req dto.CreateMasterRequest) error {
	query := fmt.Sprintf("INSERT INTO %s (name , add_price) VALUE ($1 ,$2) ", table)

	_, err := r.db.Exec(ctx, query, req.Name, req.AddPrice)

	return err
}

func (r *MasterRepository) GetAll(ctx context.Context, table string) ([]dto.Master, error) {
	query := fmt.Sprintf("SELECT id, name, add_price FROM %s", table)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[dto.Master])
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

	return err
}

func (r *MasterRepository) Delete(ctx context.Context, table string, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", table)

	_, err := r.db.Exec(ctx, query, id)
	return err
}
