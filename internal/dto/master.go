package dto

type Master struct {
	Id       int     `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	AddPrice *float64 `json:"add_price" db:"add_price"`
}

type CreateMasterRequest struct {
	Name     string  `json:"name" binding:"required"`
	AddPrice float64 `json:"add_price"`
}

type UpdateMasterRequest struct {
	Name     *string  `json:"name" binding:"required"`
	AddPrice *float64 `json:"add_price"`
}