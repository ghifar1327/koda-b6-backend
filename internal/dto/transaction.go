package dto

import (
	"github.com/google/uuid"
)

type UpdateTransactionRequst struct {
	Status string `json:"status"`
}

type CreateTransactionRequest struct {
	UserId        uuid.UUID           `json:"user_id"`
	Address       string              `json:"address"`
	IdMethod      int                 `json:"id_method"`
	PaymentMethod string              `json:"payment_method"`
	IdVoucher     *int                `json:"id_voucher"`
	Items         []CreateItemRequest `json:"items"`
}

type CreateItemRequest struct {
	ProductId int `json:"product_id"`
	SizeId    int `json:"size_id"`
	VariantId int `json:"variant_id"`
	Quantity  int `json:"quantity"`
}
