package models

type Cart struct {
	Image       string `json:"product_image" db:"product_image"`
	ProductName string `json:"product_name" db:"product_name"`
	Size        string `json:"size" db:"size"`
	Variant     string `json:"variant" db:"variant"`
	Quantity    string `json:"quantity" db:"quantity"`
	Subtotal    string `json:"subtotal" db:"subtotal"`
}