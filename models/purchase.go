package models

import "time"

type Purchase struct {
	ID              uint             `gorm:"primaryKey;autoIncrement"`
	UserID          uint             `gorm:"not null"`
	OrderDate       time.Time        `gorm:"autoCreateTime"`
	Status          string           `gorm:"type:varchar(50);not null"` // Status pesanan seperti Pending, Shipped, Delivered
	TotalAmount     float64          `gorm:"type:decimal(10,2);not null"`
	PaymentStatus   string           `gorm:"type:varchar(50);not null"` // Status pembayaran seperti Paid, Unpaid
	PurchaseDetails []PurchaseDetail `gorm:"foreignKey:OrderID"`
}

type PurchaseDetail struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	OrderID      uint    `gorm:"not null"`
	ProductID    uint    `gorm:"not null"`
	Quantity     int     `gorm:"not null"`
	PriceAtOrder float64 `gorm:"type:decimal(10,2);not null"`

	Purchase Purchase `gorm:"foreignKey:OrderID"`   // Relasi ke Order
	Product  Product  `gorm:"foreignKey:ProductID"` // Relasi ke Product
}
