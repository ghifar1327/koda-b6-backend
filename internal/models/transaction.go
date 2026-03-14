package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id               uuid.UUID    `json:"id"`
	FullName         string       `json:"full_name"`
	Address          string       `json:"address"`
	Phone            string       `json:"phone"`
	Shipping         string       `json:"shipping"`
	PaymentMethod    string       `json:"payment_method"`
	Status           string       `json:"status"`
	TotalTransaction int          `json:"total_transaction"`
	Items            []ItemDetail `json:"items"`
	CreatedAt        time.Time    `json:"created_at"`
}

type ItemDetail struct {
	TransactionId uuid.UUID `json:"transaction_id"`
	ProductImage  string    `json:"product_image"`
	ProductName   string    `json:"product_name"`
	Size          string    `json:"size"`
	Variant       string    `json:"variant"`
	Quantity      int       `json:"quantity"`
	SubTotal      int       `json:"subtotal"`
}

type TransactionRow struct {
	Id            uuid.UUID `db:"id"`
	FullName      string    `db:"full_name"`
	Address       string    `db:"address"`
	Phone         string    `db:"phone"`
	Shipping      string    `db:"shipping"`
	PaymentMethod string    `db:"payment_method"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`

	ProductName  string `db:"product_name"`
	ProductImage string `db:"product_image"`
	Size         string `db:"size"`
	Variant      string `db:"variant"`
	Quantity     int    `db:"quantity"`
	Subtotal     int    `db:"subtotal"`
}
