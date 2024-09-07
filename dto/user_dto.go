package dto

import "time"

type UserResponse struct {
	Id          int            `json:"id"`
	Email       string         `json:"email"`
	PhoneNumber string         `json:"phone_number"`
	Firstname   string         `json:"firtsname"`
	Lastname    string         `json:"lastname"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Carts       []CartResponse `json:"cart"`
	CreatedAt   time.Time      `json:"created_at"`
}
