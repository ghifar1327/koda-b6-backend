package dto

type CreateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Categories  []string `json:"categories"`
	Stock       int      `json:"stock"`
}
