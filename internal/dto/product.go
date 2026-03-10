package dto

import "time"

type CreteProductRequest struct {
	Name         string
	Description string
	Price        int
	Stoct        int
	CreatedAt    time.Time
}
