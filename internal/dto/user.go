package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUsersRequest struct {
	Picture   string   `json:"picture"`
	Email     string   `json:"email"`
	FullName  string   `json:"full_name"`
	Password  string   `json:"password"`
	Address   string   `json:"address"`
	Phone     string   `json:"phone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthResponse struct {
	Id       uuid.UUID `json:"id"`
	Picture  string    `json:"picture"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Address  string    `json:"address"`
	RoleId   int       `json:"role_id"`
	Phone    string    `json:"phone"`
}

