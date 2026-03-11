package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id             uuid.UUID `json:"id"`
	UserId         uuid.UUID `json:"user_id"`
	Status         string    `json:"status"`
	IdMethod       string    `json:"id_method"`
	PaymentMethode string    `json:"payment_method"`
	IdVoucher      int       `json:"id_voucher"`
	CreatedAt      time.Time `json:"created_at"`
}
