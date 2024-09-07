package main

import (
	"api_goshop/models"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Struct untuk tabel Product di database
type Product struct {
	ID              int                     `gorm:"primaryKey;autoIncrement"`
	Title           string                  `gorm:"type:varchar(255);not null"`
	Description     string                  `gorm:"type:varchar(255);not null"`
	Category        string                  `gorm:"type:varchar(255);not null"`
	Price           float64                 `gorm:"not null"`
	Entity          int                     `gorm:"not null"`
	PurchaseDetails []models.PurchaseDetail `gorm:"foreignKey:OrderID"`
	ProductPictures string                  `gorm:"type:varchar(255);not null"`
	CreatedAt       time.Time               `gorm:"autoCreateTime"`
}

// Struct untuk response dari API
type APIResponse struct {
	Products []struct {
		ID          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
		Stock       int     `json:"stock"`
		Thumbnail   string  `json:"thumbnail"` // URL gambar
	} `json:"products"`
}

func main2() {
	// Koneksi ke database MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/goshop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Auto-migrate untuk membuat tabel products dan pictures
	if err := db.AutoMigrate(&Product{}); err != nil {
		fmt.Println("Failed to auto-migrate:", err)
		return
	}

	// Ambil data dari API DummyJSON
	client := resty.New()
	resp, err := client.R().
		SetResult(&APIResponse{}).
		Get("https://dummyjson.com/products?limit=200")
	if err != nil {
		fmt.Println("Failed to fetch data from API:", err)
		return
	}

	// Parse response dari API
	const exchangeRate = 15000.0
	apiResponse := resp.Result().(*APIResponse)

	// Masukkan data dari API ke database dengan nilai Entity dari Stock
	for _, product := range apiResponse.Products {
		dbProduct := Product{
			ID:              product.ID,
			Title:           product.Title,
			Description:     product.Description,
			Price:           product.Price * exchangeRate, // Konversi harga ke Rupiah
			Category:        product.Category,
			Entity:          product.Stock,
			ProductPictures: product.Thumbnail}

		// Simpan produk terlebih dahulu
		if err := db.Create(&dbProduct).Error; err != nil {
			fmt.Println("Failed to create product:", err)
			continue
		}

	}

	fmt.Println("Data from API inserted successfully!")
}
