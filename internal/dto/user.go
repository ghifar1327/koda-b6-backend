package dto

type UpdateUsersRequest struct {
	Picture  string `json:"picture"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}
