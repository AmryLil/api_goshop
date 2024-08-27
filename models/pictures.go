package models

type ProductPicture struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ProductID   int    `gorm:"not null"`
	PictureData []byte `gorm:"type:blob"` // Simpan gambar sebagai BLOB
}
