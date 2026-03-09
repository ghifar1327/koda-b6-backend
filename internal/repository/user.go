package repository

import (
	"backend/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserrepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAllUser(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT 
			id, 
			full_name, 
			picture, 
			email,
			password,
			role_id, 
			phone, 
			address, 
			created_at, 
			updated_at 
			FROM users`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.Id, 
			&u.FullName, 
			&u.Picture, 
			&u.Email, 
			&u.Password, 
			&u.RoleId, 
			&u.Phone, 
			&u.Address,
			&u.CreatedAt, 
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// ==================================================================================================================================================== Get User By ID
func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
		SELECT 
		id, 
		picture, 
		full_name, 
		email, 
		password, 
		address, 
		phone, 
		role_id 
			  FROM users WHERE id=$1`

	var user models.User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.Picture,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Address,
		&user.Phone,
		&user.RoleId,
	)

	if err != nil {
		return nil, err
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

func (r *UserRepository) UpdateUser(ctx context.Context, id uuid.UUID, u models.User) error {
	query := `UPDATE users 
SET picture=$1, full_name=$2, email=$3, password=$4, address=$5, phone=$6, role_id=$7, updated_at=$8 
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
	return err
}

// ======================================================================================================== DELETE USER
func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// func (s *UserService) Register(ctx context.Context, req models.CreateUserRequest) error {
// 	if err := validateUser(req.Fullname, req.Email, req.Password); err != nil {
// 		return err
// 	}

// 	argon := argon2.DefaultConfig()
// 	encoded, err := argon.HashEncoded([]byte(req.Password))

// 	if err != nil {
// 		return err
// 	}

// 	newUser := models.User{
// 		Fullname: req.Fullname,
// 		Email:    req.Email,
// 		Password: string(encoded),
// 	}
// 	return s.repo.Create(ctx, newUser)
// }
