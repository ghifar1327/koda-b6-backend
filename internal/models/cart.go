package models

type Cart struct {
	Id          int     `json:"id" db:"id"`
	Image       string  `json:"product_image" db:"product_image"`
	ProductName string  `json:"product_name" db:"product_name"`
	Size        string  `json:"size" db:"size"`
	Variant     string  `json:"variant" db:"variant"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Subtotal    float64 `json:"subtotal" db:"subtotal"`
}
