package models

import "time"

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Profile_img string `json:"profile_img"`
	Created_at  time.Time
	Updated_at  time.Time
}

type CreateUserRequest struct {
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Profile_img string `json:"profile_img"`
	Created_at  time.Time
	Updated_at  time.Time
}