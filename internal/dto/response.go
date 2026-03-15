package dto

import "github.com/google/uuid"

type Response struct {
	Success bool   `json:"success"`
	Message string `jason:"message"`
}

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	FullName string    `json:"name"`
	Email    string    `json:"email"`
}
type ResponseToken struct {
	Success bool   `json:"success"`
	Message string `jason:"message"`
	Token   string `json:"token"`
}
