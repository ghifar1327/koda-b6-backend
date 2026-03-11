package dto

type ResetPwdRequest struct {
	Email       string `json:"email"`
	Code        int    `json:"code"`
	NewPassword string `json:"new_password"`
}
