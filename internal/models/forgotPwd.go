package models

type ForgotPassword struct {
	Email string `json:"email" db:"email"` 
	Code  int    `json:"code" db:"code"`
}
