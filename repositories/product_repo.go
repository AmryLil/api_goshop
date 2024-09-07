package repositories

import (
	"api_goshop/models"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	ReadProducts() ([]models.Product, error)
	CreateProduct(model models.Product) error
	DeleteProduct(model *models.Product, id int) error
	UpdateProduct(id int, model models.Product, key string, value interface{}) error
	SearchProducts(query string) ([]models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	GetCategories() ([]models.Product, error)
}

type product_repo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *product_repo {
	return &product_repo{db}
}

func (r product_repo) GetCategories() ([]models.Product, error) {
	var categoryProducts []models.Product
	err := r.db.Distinct("category").Find(&categoryProducts).Error
	return categoryProducts, err
}

func (r product_repo) ReadProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}
func (r product_repo) CreateProduct(model models.Product) error {
	err := r.db.Create(&model).Error
	return err
}
func (r product_repo) DeleteProduct(model *models.Product, id int) error {
	err := r.db.Delete(&model, id).Error
	return err
}
func (r product_repo) UpdateProduct(id int, model models.Product, key string, value interface{}) error {
	err := r.db.First(&model, id).Error
	if err != nil {
		return err
	}
	err = r.db.Model(&model).Update(key, value).Error
	return err
}
func (r product_repo) SearchProducts(query string) ([]models.Product, error) {
	var model []models.Product
	err := r.db.Where("title LIKE ?", "%"+query+"%").Find(&model).Error
	fmt.Println("model", model)
	return model, err
}
func (r product_repo) GetProductsByCategory(category string) ([]models.Product, error) {
	var model []models.Product
	err := r.db.Where("category = ?", category).Find(&model).Error
	return model, err
}
