package repositories

import (
	"api_goshop/models"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	CreatePayment(models models.Payment) error
	CreateItemDetailsPayment(models []models.ItemDetails) error
	FindUserByID(id *int) (models.UserAccounts, error)
}

type payment_repo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *payment_repo {
	return &payment_repo{db}
}

func (r payment_repo) CreatePayment(models models.Payment) error {
	err := r.db.Create(&models).Error
	return err
}

func (r *payment_repo) FindUserByID(id *int) (models.UserAccounts, error) {
	var user_account models.UserAccounts
	err := r.db.Find(&user_account, id).Error
	return user_account, err
}

func (r payment_repo) CreateItemDetailsPayment(models []models.ItemDetails) error {
	err := r.db.Create(&models).Error
	return err
}
