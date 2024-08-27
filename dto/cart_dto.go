package dto

import "time"

// request

type CartRequest struct {
	ID     int `json:"id" `
	UserID int `json:"user_id" `
}

type CartItemRequest struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// response
type CartResponse struct {
	ID        int                `json:"id"`
	UserID    int                `json:"user_id"`
	CartItems []CartItemResponse `json:"cart_items"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CartItemResponse struct {
	ID       int             `json:"id"`
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
}
