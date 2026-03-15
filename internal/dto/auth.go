package dto

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ResetPwdRequest struct {
	Email       string `json:"email"`
	Code        int    `json:"code"`
	NewPassword string `json:"new_password"`
}
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}
