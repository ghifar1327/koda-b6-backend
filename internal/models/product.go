package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
	Categories  []string  `json:"categories" db:"categories"`
	Stock       int       `json:"stock" db:"stock"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ReviewProduct struct {
	Id                  int       `json:"id" db:"id"`
	UserId              uuid.UUID `json:"user_id" db:"user_id"`
	IdTransactionDetail int       `json:"id_transaction_details" db:"id_transaction_details"`
	Rating              float64   `json:"rating" db:"rating"`
	Message             string    `json:"message" db:"message"`
}

type Reviews struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Image       string `json:"image" db:"images"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	TotalReview int    `json:"total_review" db:"total_review"`
}

type RecommendedProduct struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Image       string  `json:"image" db:"images"`
	Description string  `json:"description" db:"description"`
	Price       int     `json:"price" db:"price"`
	TotalReview int     `json:"total_review" db:"total_review"`
	AvgRating   float64 `json:"avg_rating" db:"avg_rating"`
}