package dto

import (
	"github.com/google/uuid"
)

type UpdateUsersRequest struct {
	Picture  string `json:"picture"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	RoleId   int    `json:"role_id"`
}

type userResponse struct {
	Id       uuid.UUID `json:"id"`
	Picture  string    `json:"picture"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Address  string    `json:"address"`
	RoleId   int       `json:"role_id"`
	Phone    string    `json:"phone"`
}
