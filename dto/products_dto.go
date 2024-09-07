package dto

type ProductResponse struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Entity          int     `json:"entity"`
	Category        string  `json:"category"`
	ProductPictures string  `json:"product_picture"`
}
type ProductRequest struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Entity          int     `json:"entity"`
	Category        string  `json:"category"`
	ProductPictures string  `json:"product_picture"`
}

type ProductIDRequest struct {
	ID int `json:"id" binding:"required"`
}
