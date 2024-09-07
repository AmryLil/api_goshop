package handler

import (
	"api_goshop/config"
	"api_goshop/dto"
	"api_goshop/handleError"
	"api_goshop/helper"
	"api_goshop/services"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type products_handler struct {
	service services.ProductServices
}

func NewProductsHandler(service services.ProductServices) *products_handler {
	return &products_handler{service}
}

func (h products_handler) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	products, err := h.service.GetProductsByCategory(category)
	if err != nil {
		handleError.HandleError(c, err)
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Get products by category successfull",
		Data:       products,
	})
	c.JSON(http.StatusOK, res)
}

func (h products_handler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories()
	if err != nil {
		handleError.HandleError(c, err)
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Get Categories is Succesfull",
		Data:       categories,
	})
	c.JSON(http.StatusOK, res)
}

func (h products_handler) GetAllProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		products, err := h.service.GetAllProduct()
		if err != nil {
			handleError.HandleError(c, err)
		}
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "Get data succesfully!!!",
			Data:       products,
		})
		c.JSON(http.StatusOK, res)
	} else {
		results, err := h.service.SearchProducts(query)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to search product"})
			return
		}
		res := helper.Response(dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "Search succesfully",
			Data:       results,
		})
		c.JSON(http.StatusOK, res)
	}

}

func (h products_handler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var input map[string]any
	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update product"})
		return
	}

	for key, value := range input {
		fmt.Println(key, value)
		if err := h.service.UpdateProduct(id, key, value); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Product succesfully updated",
	})
	c.JSON(http.StatusOK, res)
}

func (h products_handler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.service.DeleteProduct(id)
	if err != nil {
		handleError.HandleError(c, err)
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Product Succesfully Deleted",
	})
	c.JSON(http.StatusOK, res)
}

func (h products_handler) CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	price := c.PostForm("price")
	entitystr := c.PostForm("entity")
	category := c.PostForm("category")

	// Validasi input produk
	if name == "" || description == "" || price == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product data"})
		return
	}
	fmt.Println("1")

	// Konversi harga ke float64
	priceValue, err := strconv.ParseFloat(price, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}
	fmt.Println("2")

	entity, err := strconv.Atoi(entitystr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity"})
		return
	}
	fmt.Println("3")

	client := config.InitFirebase()

	// Ambil foto dari form
	product_picture, err := c.FormFile("product_picture")
	if err != nil {
		fmt.Println("Error getting file from request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	fmt.Println("4")

	src, err := product_picture.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to open the file"})
		return
	}
	defer src.Close()

	// Mengunggah file ke Firebase
	ctx := context.Background()
	bucketName := "goshop-be94c.appspot.com" // Ganti dengan nama bucket Anda
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get storage bucket"})
		return
	}

	// Upload file ke Firebase Storage
	fmt.Println("5")
	obj := bucket.Object(product_picture.Filename)
	wc := obj.NewWriter(ctx)
	if _, err = io.Copy(wc, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}
	if err := wc.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close file writer"})
		return
	}
	fmt.Println("6")

	// Mendapatkan URL unduhan
	downloadURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucketName, url.PathEscape(product_picture.Filename))

	reqProduct := dto.ProductRequest{
		Name:            name,
		Description:     description,
		Price:           priceValue,
		Entity:          entity,
		Category:        category,
		ProductPictures: downloadURL,
	}

	err = h.service.CreateProduct(&reqProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Product successfully created!!!",
	})
	c.JSON(http.StatusCreated, res)
}
