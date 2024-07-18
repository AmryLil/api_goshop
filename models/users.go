package models

import "time"

type UserAccounts struct {
	Id              int    `gorm:"primaryKey;autoIncrement"`
	Email           string `gorm:"type:varchar(255);not null"`
	PhoneNumber     string `gorm:"type:varchar(255);"`
	Firstname       string `gorm:"type:varchar(255);not null"`
	Lastname        string `gorm:"type:varchar(255);not null"`
	Username        string `gorm:"type:varchar(255);not null"`
	Password        string `gorm:"type:varchar(255);not null"`
	ProfilePicture  []byte
	IsSeller        bool            `gorm:"type:boolean;not null;default:false"`
	Seller          Seller          `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Purchase        Purchase        `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PurchaseHistory PurchaseHistory `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Cart            Cart            `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt       time.Time       `gorm:"autoCreateTime"`
}

type Seller struct {
	Id             int       `gorm:"primaryKey;autoIncrement"`
	UserId         int       `gorm:"unique;not null"`
	StoreName      string    `gorm:"type:varchar(255);not null"`
	Products_Id    []Product `gorm:"foreignKey:SellerId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Description    string    `gorm:"type:text"`
	FullName       string    `gorm:"type:varchar(255);not null"`
	Email          string    `gorm:"type:varchar(255);not null"`
	PhoneNumber    string    `gorm:"type:varchar(255);not null"`
	IdentityNumber string    `gorm:"type:varchar(255);not null"`
}

type Cart struct {
	Id             int    `gorm:"primaryKey;autoIncrement"`
	UserId         int    `gorm:"unique;not null"`
	ProductName    string `gorm:"type:varchar(255);not null"`
	Entity         int    `gorm:"not null"`
	Type           string `gorm:"type:varchar(255);not null"`
	Variant        string `gorm:"type:varchar(255);not null"`
	StoreName      string `gorm:"type:varchar(255);not null"`
	ProductPicture []byte
}
