package dto

import "github.com/google/uuid"

type UpdateReviewProductRequest struct {
	Rating  float64 `json:"rating"`
	Message string  `json:"message"`
}
type CreateReviewProductRequest struct {
	UserId              uuid.UUID `json:"user_id"`
	IdTransactionDetail int       `json:"id_transaction_details"`
	Rating              float64   `json:"rating"`
	Message             string    `json:"message"`
}
