package dto

import "github.com/google/uuid"

type UpdateTransactionRequst struct {
	Status string `json:"status"`
}
type CreateRransactionRequest struct {
	Id            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	IdMethod      int       `json:"id_method"`
	IdVoucher     *int      `json:"id_voucher"`
}

type AddItemRequiest struct {
	TransactionId uuid.UUID `json:"transaction_id"`
	ProductId     int       `json:"product_id"`
	SizeId        int       `json:"size_id"`
	VariantId     int       `json:"variant_id"`
	Quantity      int       `json:"quantity"`
}
