package repositories

import (
	"api_goshop/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	ReadProducts() ([]models.Product, error)
}

type product_repo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *product_repo {
	return &product_repo{db}
}
