package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id               uuid.UUID    `json:"id" db:"id"`
	FullName         string       `json:"full_name" db:"full_name"`
	Address          string       `json:"address" db:"address"`
	Phone            string       `json:"phone" db:"phone"`
	Shipping         string       `json:"shipping" db:"shipping"`
	PaymentMethod    string       `json:"payment_method" db:"payment_method"`
	Status           string       `json:"status" db:"status"`
	TotalTransaction int          `json:"total_transaction" db:"total_transaction"`
	Items            []ItemDetail `json:"items" db:"items"`
	CreatedAt        time.Time    `json:"created_at" db:"created_at"`
}

type ItemDetail struct {
	TransactionId uuid.UUID `json:"transaction_id" db:"transaction_id"`
	ProductImage  string    `json:"product_image" db:"product_image"`
	ProductName   string    `json:"product_name" db:"product_name"`
	Size          string    `json:"size" db:"size"`
	Variant       string    `json:"variant" db:"variant"`
	Quantity      int       `json:"quantity" db:"quantity"`
	SubTotal      int       `json:"subtotal" db:"subtotal"`
}

type TransactionRow struct {
	Id            uuid.UUID `json:"id" db:"id"`
	FullName      string    `json:"full_name" db:"full_name"`
	Address       string    `json:"address" db:"address"`
	Phone         string    `json:"phone" db:"phone"`
	Shipping      string    `json:"shipping" db:"shipping"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`

	ProductName  string `json:"product_name" db:"product_name"`
	ProductImage string `json:"product_image" db:"product_image"`
	Size         string `json:"size" db:"size"`
	Variant      string `json:"variant" db:"variant"`
	Quantity     int    `json:"quantity" db:"quantity"`
	Subtotal     int    `json:"subtotal" db:"subtotal"`
}
