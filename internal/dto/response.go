package dto

import "github.com/google/uuid"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Picture  string    `json:"picture"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Role     string    `json:"role"`
}
type ResponseToken struct {
	Success bool         `json:"success"`
	Message string       `jason:"message"`
	Token   string       `json:"token"`
	User    UserResponse `json:"user"`
}

type ResponseWrap struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}
