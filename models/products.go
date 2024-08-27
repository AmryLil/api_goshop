package models

import "time"

type Product struct {
	ID              int              `gorm:"primaryKey;autoIncrement"`
	Title           string           `gorm:"type:varchar(255);not null"`
	Description     string           `gorm:"type:varchar(255);not null"`
	Category        string           `gorm:"type:varchar(255);not null"`
	Entity          int              `gorm:"not null"`
	PurchaseDetails []PurchaseDetail `gorm:"foreignKey:OrderID"`
	CartItems       []CartItem       `gorm:"foreignKey:ProductID"`
	ProductPictures []ProductPicture `gorm:"foreignKey:ProductID"`
	CreatedAt       time.Time        `gorm:"autoCreateTime"`
}
