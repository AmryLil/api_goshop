package models

import "time"

type Payment struct {
	ID          int           `gorm:"primaryKey;autoIncrement"`
	Amount      float64       `gorm:"not null"`
	UserID      int           `gorm:"not null"`
	Status      string        `gorm:"not null"`
	ItemDetails []ItemDetails `gorm:"foreignKey:PaymentID"`
	CreatedAt   time.Time     `gorm:"autoCreateTime"`
}

type ItemDetails struct {
	ProudctID int       `gorm:"not null"`
	PaymentID int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	QTY       int       `gorm:"not null"`
	Price     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
