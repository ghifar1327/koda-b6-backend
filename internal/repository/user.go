package repository

import (
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

type UserRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewUserrepository(db *pgxpool.Pool, rdb *redis.Client) *UserRepository {
	return &UserRepository{
		db:  db,
		rdb: rdb,
	}
}

func (r *UserRepository) GetAllUser(ctx context.Context) ([]models.User, error) {
	key := "users"

	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result []models.User
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result, nil
		}
	}

	query := `
		SELECT 
			id, 
			full_name, 
			COALESCE(picture, '') AS picture, 
			email,
			password,
			role_id, 
			phone, 
			address, 
			created_at, 
			updated_at FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(users)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}

	return users, nil
}

// ==================================================================================================================================================== Get User By ID
func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	fmt.Println(id)
	key := fmt.Sprintf("user:id:%s", id.String())

	// ================= CACHE =================
	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result models.User
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}

	// ================= DB =================
	query := `
		SELECT 
			id, 
			COALESCE(picture, '') AS picture, 
			full_name, 
			email, 
			password, 
			address, 
			phone, 
			role_id,
			created_at,
			updated_at
		FROM users 
		WHERE id=$1
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	// ================= CACHE =================
	data, err := json.Marshal(user)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}

	return &user, nil
}

// ====================================================================================================================================================  Create User

func (r *UserRepository) CreateUser(ctx context.Context, u models.User) error {
	query := `INSERT INTO users (id, full_name, email, password, address, phone, role_id,created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(ctx, query,
		u.Id,
		u.FullName,
		u.Email,
		u.Password,
		u.Address,
		u.Phone,
		u.RoleId,
		u.CreatedAt)

	return err
}

// ==================================================================================================================================================== Update User

func (r *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, u models.User) (models.User, error) {
	query := `
	    UPDATE users 
	    SET 
	    	picture=$1, 
	    	full_name=$2, 
	    	email=$3, 
	    	password=$4, 
	    	address=$5, 
	    	phone=$6, 
	    	role_id=$7,
	    	updated_at=$8 
	    WHERE id=$9`
	_, err := r.db.Exec(ctx, query,
		u.Picture,
		u.FullName,
		u.Email,
		u.Password,
		u.Address,
		u.Phone,
		u.RoleId,
		time.Now(),
		id)
	if err != nil {
		return models.User{}, err
	}
	r.rdb.Del(ctx, fmt.Sprintf("user:id:%s", id.String()))
	r.rdb.Del(ctx, fmt.Sprintf("user:email:%s", u.Email))
	user, err := r.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

// ======================================================================================================== DELETE USER
func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	r.rdb.Del(ctx, fmt.Sprintf("user:id:%s", id.String()))
	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	key := fmt.Sprintf("user:email:%s", email)

	cached, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		var result models.User
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}
	query := `
		SELECT 
			id,
			full_name,
			COALESCE(picture, '') AS picture, 
			email,
			password,
			role_id,
			phone,
			address,
			created_at,
			updated_at
		FROM users
		WHERE email=$1
	`
	row, err := r.db.Query(ctx, query, email)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(user)
	if err == nil {
		r.rdb.Set(ctx, key, data, time.Minute*15)
	}

	return &user, nil
}
