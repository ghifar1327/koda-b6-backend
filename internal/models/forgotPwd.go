package models

type ForgotPassword struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}
