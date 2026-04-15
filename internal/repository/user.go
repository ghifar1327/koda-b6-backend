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
     		u.id, 
     		u.full_name, 
     		COALESCE(u.picture, '') AS picture, 
     		u.email,
     		u.password,
     		r.name as role, 
     		u.phone, 
     		u.address, 
     		u.created_at, 
     		u.updated_at FROM users u
        JOIN roles r ON u.role_id = r.id;`
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
     		u.id, 
     		u.full_name, 
     		COALESCE(u.picture, '') AS picture, 
     		u.email,
     		u.password,
     		r.name as role, 
     		u.phone, 
     		u.address, 
     		u.created_at, 
     		u.updated_at FROM users u
        JOIN roles r ON u.role_id = r.id
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
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// get role_id from role name
	var roleID int
	err = tx.QueryRow(ctx,
		`SELECT id FROM roles WHERE name = $1`,
		u.Role,
	).Scan(&roleID)

	if err != nil {
		return fmt.Errorf("role '%s' not found: %w", u.Role, err)
	}

	// insert nuew user
	query := `
		INSERT INTO users 
		(id, full_name, email, password, address, phone, role_id, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = tx.Exec(ctx, query,
		u.Id,
		u.FullName,
		u.Email,
		u.Password,
		u.Address,
		u.Phone,
		roleID,
		u.CreatedAt,
	)

	if err != nil {
		return err
	}

	// 
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	r.rdb.Del(ctx, "users")
	return nil
}

// ==================================================================================================================================================== Update User

func (r *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, u models.User) (models.User, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return models.User{}, err
	}
	defer tx.Rollback(ctx)

	// get role_id from role name
	var roleID int
	err = tx.QueryRow(ctx,
		`SELECT id FROM roles WHERE name = $1`,
		u.Role,
	).Scan(&roleID)

	if err != nil {
		return models.User{}, fmt.Errorf("role '%s' not found: %w", u.Role, err)
	}

	// update user
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
	    WHERE id=$9
	`
	_, err = tx.Exec(ctx, query,
		u.Picture,
		u.FullName,
		u.Email,
		u.Password,
		u.Address,
		u.Phone,
		roleID,
		time.Now(),
		id,
	)

	if err != nil {
		return models.User{}, err
	}

	// commit
	if err := tx.Commit(ctx); err != nil {
		return models.User{}, err
	}

	// invalidate cache
	r.rdb.Del(ctx, fmt.Sprintf("user:id:%s", id.String()))
	r.rdb.Del(ctx, fmt.Sprintf("user:email:%s", u.Email))

	// get updated user
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
     		u.id, 
     		u.full_name, 
     		COALESCE(u.picture, '') AS picture, 
     		u.email,
     		u.password,
     		r.name as role, 
     		u.phone, 
     		u.address, 
     		u.created_at, 
     		u.updated_at FROM users u
        JOIN roles r ON u.role_id = r.id
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
