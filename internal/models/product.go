package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int       `db:"price"`
	Categories  []string  `db:"categories"`
	Stock       int       `db:"stock"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ReviewProduct struct {
	Id                  int       `db:"id"`
	UserId              uuid.UUID `db:"user_id"`
	IdTransactionDetail int       `db:"id_transaction_details"`
	Rating              float64   `db:"rating"`
	Message             string    `db:"message"`
}

type Reviews struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Image       string `db:"images"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	TotalReview int    `db:"total_review"`
}

type RecommendedProduct struct {
	Id          int     `db:"id"`
	Name        string  `db:"name"`
	Image       string  `db:"images"`
	Description string  `db:"description"`
	Price       int     `db:"price"`
	TotalReview int     `db:"total_review"`
	AvgRating   float64 `db:"avg_rating"`
}
