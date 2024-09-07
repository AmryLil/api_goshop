package models

import (
	"time"
)

type Product struct {
	ID              int              `gorm:"primaryKey;autoIncrement"`
	Title           string           `gorm:"type:varchar(255);not null"`
	Description     string           `gorm:"type:varchar(255);not null"`
	Category        string           `gorm:"type:varchar(255);not null"`
	Price           float64          `gorm:"not null"`
	Entity          int              `gorm:"not null"`
	PurchaseDetails []PurchaseDetail `gorm:"foreignKey:OrderID"`
	CartItems       []CartItem       `gorm:"foreignKey:ProductID"`
	ProductPictures string           `gorm:"type:varchar(255);not null"`
	CreatedAt       time.Time        `gorm:"autoCreateTime"`
}
