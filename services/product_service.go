package services

import (
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/models"
	"api_goshop/repositories"
)

type ProductServices interface {
	GetAllProduct() ([]dto.ProductResponse, error)
	CreateProduct(req *dto.ProductRequest) error
	DeleteProduct(id int) error
	UpdateProduct(id int, key string, value any) error
	SearchProducts(query string) ([]dto.ProductResponse, error)
	GetCategories() ([]any, error)
	GetProductsByCategory(category string) ([]dto.ProductResponse, error)
}

type product_service struct {
	repo repositories.ProductRepository
}

func NewProductsService(repo repositories.ProductRepository) *product_service {
	return &product_service{repo}
}

func (s product_service) GetProductsByCategory(category string) ([]dto.ProductResponse, error) {
	var productResponses []dto.ProductResponse
	products, err := s.repo.GetProductsByCategory(category)
	if err != nil {
		return nil, &handleError.InternalServerError{Message: err.Error()}
	}
	for _, product := range products {
		productResponse := dto.ProductResponse{
			ID:              product.ID,
			Name:            product.Title,
			Description:     product.Description,
			Price:           product.Price,
			Entity:          product.Entity,
			Category:        product.Category,
			ProductPictures: product.ProductPictures,
		}
		productResponses = append(productResponses, productResponse)

	}
	return productResponses, err
}

func (s product_service) GetCategories() ([]any, error) {
	var categoryProducts []any
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, &handleError.InternalServerError{Message: err.Error()}
	}
	for _, category := range categories {
		categoryProducts = append(categoryProducts, category.Category)
	}
	return categoryProducts, err
}

func (s product_service) SearchProducts(query string) ([]dto.ProductResponse, error) {
	var productResponses []dto.ProductResponse
	products, err := s.repo.SearchProducts(query)
	if err != nil {
		return nil, &handleError.InternalServerError{Message: err.Error()}
	}
	for _, product := range products {
		productResponse := dto.ProductResponse{
			ID:              product.ID,
			Name:            product.Title,
			Description:     product.Description,
			Price:           product.Price,
			Entity:          product.Entity,
			Category:        product.Category,
			ProductPictures: product.ProductPictures,
		}
		productResponses = append(productResponses, productResponse)

	}
	return productResponses, err

}

func (s product_service) UpdateProduct(id int, key string, value any) error {
	var model models.Product
	if err := s.repo.UpdateProduct(id, model, key, value); err != nil {
		return err
	}
	return nil
}
func (s product_service) GetAllProduct() ([]dto.ProductResponse, error) {
	var productResponses []dto.ProductResponse
	products, err := s.repo.ReadProducts()
	if err != nil {
		return nil, &handleError.NotFoundError{Message: err.Error()}
	}
	for _, product := range products {
		productResponse := dto.ProductResponse{
			ID:              product.ID,
			Name:            product.Title,
			Description:     product.Description,
			Price:           product.Price,
			Entity:          product.Entity,
			Category:        product.Category,
			ProductPictures: product.ProductPictures,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil

}

func (s product_service) CreateProduct(req *dto.ProductRequest) error {
	model := models.Product{
		Title:           req.Name,
		Description:     req.Description,
		Category:        req.Category,
		Price:           req.Price,
		Entity:          req.Entity,
		ProductPictures: req.ProductPictures,
	}
	err := s.repo.CreateProduct(model)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s product_service) DeleteProduct(id int) error {
	var model *models.Product
	err := s.repo.DeleteProduct(model, id)
	return err
}
