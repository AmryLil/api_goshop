package services

import (
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/models"
	"api_goshop/repositories"
)

type CartService interface {
	AddtoCart(req dto.CartRequest) error
	Delete(id int, dataProduct models.Cart) error
	Update(req dto.CartRequest) error
	ReadCart() ([]models.Cart, error)
}

type cart_service struct {
	repository repositories.CartRepository
}

func NewCartService(repository repositories.CartRepository) *cart_service {
	return &cart_service{repository}
}

func (s *cart_service) AddtoCart(req dto.CartRequest) error {
	dataProduct := models.Cart{
		UserId:         req.UserId,
		ProductName:    req.ProductName,
		Entity:         req.Entity,
		Type:           req.Type,
		Variant:        req.Variant,
		StoreName:      req.StoreName,
		ProductPicture: req.ProductPicture,
	}
	err := s.repository.AddtoCart(dataProduct)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *cart_service) Update(req dto.CartRequest) error {
	var dataProduct = models.Cart{
		Id:             req.Id,
		UserId:         req.UserId,
		ProductName:    req.ProductName,
		Entity:         req.Entity,
		Type:           req.Type,
		Variant:        req.Variant,
		StoreName:      req.StoreName,
		ProductPicture: req.ProductPicture,
	}
	err := s.repository.Update(dataProduct)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *cart_service) Delete(id int, dataProdust models.Cart) error {
	err := s.repository.Delete(id, dataProdust)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *cart_service) ReadCart() ([]models.Cart, error) {
	data, err := s.repository.ReadCart()
	return data, err
}
