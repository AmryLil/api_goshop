package models

import "time"

type UserAccounts struct {
	Id          int        `gorm:"primaryKey;autoIncrement"`
	Email       string     `gorm:"type:varchar(255);not null;unique"`
	PhoneNumber string     `gorm:"type:varchar(255);"`
	Firstname   string     `gorm:"type:varchar(255);not null"`
	Lastname    string     `gorm:"type:varchar(255);not null"`
	Username    string     `gorm:"type:varchar(255);not null;unique"`
	Password    string     `gorm:"type:varchar(255);not null"`
	Purchases   []Purchase `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Carts       []Cart     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
}
