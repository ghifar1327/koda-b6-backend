package repository

import (
	"backend/internal/models"
	"context"

	"github.com/gin-gonic/gin"
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

func (r *UserRepository) GetAll(ctx *gin.Context) ([]models.User, error) {
	query := `SELECT id, full_name, picture, email,password ,role_id ,phone, address, created_at, updated_at FROM users`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.FullName, &u.Email, &u.Password, &u.Picture, &u.Address, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Create User
func (r *UserRepository) Create(ctx context.Context, u models.User) error {
	query := `INSERT INTO users (id, full_name, email, password, address, phone, role_id,created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(ctx, query, u.Id, u.FullName, u.Email, u.Password, u.Address, u.Phone, u.RoleId ,u.CreatedAt)
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
