package models

import "time"

// type Cart struct {
// 	Id             int    `gorm:"primaryKey;autoIncrement"`
// 	UserId         int    `gorm:"not null"`
// 	ProductName    string `gorm:"type:varchar(255);not null"`
// 	Entity         int    `gorm:"not null"`
// 	Type           string `gorm:"type:varchar(255);not null"`
// 	Price          string `gorm:"type:varchar(255);not null"`
// 	Variant        string `gorm:"type:varchar(255);not null"`
// 	StoreName      string `gorm:"type:varchar(255);not null"`
// 	ProductPicture []byte
// }

type Cart struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	UserID    int        `gorm:"not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	CartItems []CartItem `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"` // Relasi satu-ke-banyak dengan CartItem
}

type CartItem struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	CartID    int     `gorm:"not null"`
	ProductID int     `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"` // Relasi ke Product
}
