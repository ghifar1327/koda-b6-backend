package dto

type UpdateReviewProductRequest struct {
	Rating  float64 `json:"rating"`
	Message string  `json:"message"`
}
