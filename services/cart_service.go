package services

import (
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/models"
	"api_goshop/repositories"
	"errors"

	"gorm.io/gorm"
)

type CartService interface {
	AddtoCart(req dto.CartItemRequest, id *int) error
	Delete(id int, userID *int) error
	Update(req dto.CartRequest) error
	ReadCart(userID *int) (models.Cart, error)
}

type cart_service struct {
	repository repositories.CartRepository
}

func NewCartService(repository repositories.CartRepository) *cart_service {
	return &cart_service{repository}
}

func (s *cart_service) AddtoCart(req dto.CartItemRequest, userID *int) error {
	// Temukan atau buat Cart untuk user
	cart, err := s.repository.FindUserOrCreate(userID)
	if err != nil {
		return &handleError.InternalServerError{Message: "Failed to find or create cart: " + err.Error()}
	}

	// Periksa apakah produk sudah ada di dalam cart
	cartItem, err := s.repository.CartExist(int(req.ProductID), *userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &handleError.InternalServerError{Message: "Failed to check cart item existence: " + err.Error()}
	}

	if cartItem != nil {
		// Produk sudah ada, tambahkan quantity
		err = s.repository.AddQty(cartItem.ID, req.Quantity)
		if err != nil {
			return &handleError.InternalServerError{Message: "Failed to add quantity: " + err.Error()}
		}
	} else {
		var newCartItem = models.CartItem{
			CartID:    cart.ID,
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
		}
		err := s.repository.CreateCartItems(newCartItem)
		if err != nil {
			return &handleError.InternalServerError{Message: "Failed to create cart item: " + err.Error()}
		}
	}

	return nil
}

func (s *cart_service) Update(req dto.CartRequest) error {
	var dataProduct = models.Cart{
		ID:     req.ID,
		UserID: req.UserID,
	}
	err := s.repository.Update(dataProduct)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *cart_service) Delete(id int, userID *int) error {
	var cartItem models.Cart
	err := s.repository.Delete(id, userID, cartItem)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *cart_service) ReadCart(userID *int) (models.Cart, error) {
	data, err := s.repository.ReadCart(userID)
	return data, err
}
