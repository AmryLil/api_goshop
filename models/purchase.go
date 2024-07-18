package models

type Purchase struct {
	Id             int    `gorm:"primaryKey;autoIncrement"`
	SellerId       int    `gorm:"unique;not null"`
	UserId         int    `gorm:"unique;not null"`
	ProductId      int    `gorm:"unique;not null"`
	ProductName    string `gorm:"type:varchar(255);not null"`
	Entity         int    `gorm:"not null"`
	Type           string `gorm:"type:varchar(255);not null"`
	Status         bool   `gorm:"type:boolean;not null;default:false"`
	Variant        string `gorm:"type:varchar(255);not null"`
	StoreName      string `gorm:"type:varchar(255);not null"`
	ProductPicture []byte
}

type PurchaseHistory struct {
	Id             int    `gorm:"primaryKey;autoIncrement"`
	SellerId       int    `gorm:"unique;not null"`
	UserId         int    `gorm:"unique;not null"`
	ProductId      int    `gorm:"unique;not null"`
	ProductName    string `gorm:"type:varchar(255);not null"`
	Entity         int    `gorm:"not null"`
	Type           string `gorm:"type:varchar(255);not null"`
	Variant        string `gorm:"type:varchar(255);not null"`
	StoreName      string `gorm:"type:varchar(255);not null"`
	ProductPicture []byte
}
