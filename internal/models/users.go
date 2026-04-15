package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Picture   string    `json:"picture" db:"picture"`
	FullName  string    `json:"full_name" db:"full_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Address   string    `json:"address" db:"address"`
	Phone     string    `json:"phone" db:"phone"`
	Role   string       `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}