package dto

type CheckoutRequest struct {
	UserId   int    `json:"user_id"`
	Address  string `json:"address"`
	Delivery int    `json:"delivery"`
}