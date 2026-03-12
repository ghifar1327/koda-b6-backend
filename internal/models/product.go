package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Categories  []string  `json:"categories"`
	Stoct       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UploadedAt  time.Time `json:"updated_at"`
}

type ReviewProduct struct {
	Id                  int       `db:"id"`
	UserId              uuid.UUID `db:"user_id"`
	IdTransactionDetail int       `db:"id_transaksion_detail"`
	Rating              float64   `db:"rating"`
	Message             string    `db:"message"`
}

type Reviews struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Image       string `db:"url"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	TotalReview int    `db:"total_review"`
}

type RecommendedProduct struct {
	Id          int     `db:"id"`
	Name        string  `db:"name"`
	Image       string  `db:"url"`
	Description string  `db:"description"`
	Price       int     `db:"price"`
	TotalReview int     `db:"total_review"`
	AvgRating   float64 `db:"avg_rating"`
}
