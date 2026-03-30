package dto

import "github.com/google/uuid"

type ADDCartRequest struct {
	Id        int       `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	ProductID int       `json:"product_id" db:"product_id"`
	SizeID    int       `json:"size_id" db:"size_id"`
	Variant   int       `json:"variant_id" db:"variant_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
}
