package models

type ForgotPassword struct {
	Id    int    `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Code  string `json:"code" db:"code"`
}