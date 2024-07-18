package models

import "time"

type Product struct {
	Id              int             `gorm:"primaryKey;autoIncrement"`
	ProductName     string          `gorm:"type:varchar(255);not null"`
	Entity          int             `gorm:"not null"`
	Type            string          `gorm:"type:varchar(255);not null"`
	Purchase        Purchase        `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PurchaseHistory PurchaseHistory `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SellerId        int             `gorm:"unique;not null"`
	Variant         string          `gorm:"type:varchar(255);not null"`
	StoreName       string          `gorm:"type:varchar(255);not null"`
	ProductPicture  []byte
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
