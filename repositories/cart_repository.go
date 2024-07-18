package repositories

import (
	"api_goshop/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddtoCart(dataProduct models.Cart) error
	Delete(id int, dataProduct models.Cart) error
	Update(dataProduct models.Cart) error
	ReadCart() ([]models.Cart, error)
}

type cart_repository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cart_repository {
	return &cart_repository{db}
}

func (r cart_repository) AddtoCart(dataProduct models.Cart) error {
	err := r.db.Create(&dataProduct).Error
	return err
}

func (r cart_repository) Delete(id int, dataProduct models.Cart) error {
	err := r.db.Delete(&dataProduct, id).Error
	return err
}

func (r cart_repository) Update(dataProduct models.Cart) error {
	err := r.db.Save(&dataProduct).Error
	return err
}

func (r cart_repository) ReadCart() ([]models.Cart, error) {
	var dataProduct []models.Cart
	err := r.db.Find(&dataProduct).Error
	return dataProduct, err
}
