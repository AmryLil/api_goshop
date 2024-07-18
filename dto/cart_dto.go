package dto

type CartRequest struct {
	Id             int    `json:"id" binding:"required"`
	UserId         int    `json:"user_id" binding:"required"`
	ProductName    string `json:"product_name" binding:"required"`
	Entity         int    `json:"entity" binding:"required"`
	Type           string `json:"type" binding:"required"`
	Variant        string `json:"variant" binding:"required"`
	StoreName      string `json:"store_name" binding:"required"`
	ProductPicture []byte `json:"product_picture" binding:"required"`
}

type CartResponse struct {
	UserId         int    `json:"user_id"`
	ProductName    string `json:"product_name"`
	Entity         int    `json:"entity"`
	Type           string `json:"type"`
	Variant        string `json:"variant"`
	StoreName      string `json:"store_name"`
	ProductPicture []byte `json:"product_picture"`
}
