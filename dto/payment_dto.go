package dto

import "time"

type PaymentRequest struct {
	Amount float64 `json:"amount"`
	UserID int     `json:"user_id"`
	Status string  `json:"status"`
}
type ItemDetailsRequest struct {
	PaymentID int    `json:"payment_id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	QTY       int    `json:"qty"`
	Price     int    `json:"price"`
}
type PaymentResponse struct {
	Amount          float64               `json:"amount"`
	UserID          int                   `json:"user_id"`
	Status          string                `json:"status"`
	PaymentToken    string                `json:"payment_token"`
	PaymentMethod   string                `json:"payment_method"`
	CustomerDetails UserResponse          `json:"customer_details"`
	ItemDetails     []ItemDetailsResponse `json:"item_details"`
	CreatedAt       time.Time             `json:"created_at"`
}
type ItemDetailsResponse struct {
	PaymentID int       `json:"payment_id"`
	ProductID int       `json:"product_id"`
	Name      string    `json:"name"`
	QTY       int       `json:"qty"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
