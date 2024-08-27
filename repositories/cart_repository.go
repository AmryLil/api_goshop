package repositories

import (
	"api_goshop/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddtoCart(dataProduct models.Cart) error
	Delete(id int, userID *int, dataProduct models.Cart) error
	Update(dataProduct models.Cart) error
	ReadCart(userID *int) ([]models.Cart, error)
	FindUserOrCreate(userID *int) (models.Cart, error)
	CartExist(productID int, userID int) (*models.CartItem, error)
	CreateCartItems(cartItem models.CartItem) error
	AddQty(cartItemID int, quantity int) error
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

func (r cart_repository) Delete(id int, userID *int, dataProduct models.Cart) error {
	err := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&dataProduct).Error
	return err
}

func (r cart_repository) Update(dataProduct models.Cart) error {
	err := r.db.Save(&dataProduct).Error
	return err
}

func (r cart_repository) ReadCart(userID *int) ([]models.Cart, error) {
	var dataProduct []models.Cart
	err := r.db.Preload("CartItems.Product").First(&dataProduct, userID).Error
	return dataProduct, err
}
func (r cart_repository) FindUserOrCreate(userID *int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ?", *userID).FirstOrCreate(&cart, models.Cart{UserID: *userID}).Error
	return cart, err
}
func (r cart_repository) CartExist(productID int, userID int) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("cart_items.product_id = ? AND carts.user_id = ?", productID, userID).
		First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}
func (r cart_repository) CreateCartItems(cartItem models.CartItem) error {
	return r.db.Create(&cartItem).Error
}
func (r cart_repository) AddQty(cartItemID int, quantity int) error {
	return r.db.Model(&models.CartItem{}).Where("id = ?", cartItemID).
		Update("quantity", gorm.Expr("quantity + ?", quantity)).Error
}
