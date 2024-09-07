package routes

import (
	"api_goshop/config"
	"api_goshop/handler"
	"api_goshop/repositories"
	"api_goshop/services"

	"github.com/gin-gonic/gin"
)

func ProductRouters(c *gin.RouterGroup) {
	productRepo := repositories.NewProductRepo(config.DB)
	productService := services.NewProductsService(productRepo)
	productHandler := handler.NewProductsHandler(productService)
	c.GET("/products", productHandler.GetAllProducts)
	c.GET("/categories", productHandler.GetCategories)
	c.GET("/categories/:category", productHandler.GetProductsByCategory)
	c.POST("/create_product", productHandler.CreateProduct)
	c.DELETE("/delete_product/:id", productHandler.DeleteProduct)
	c.PATCH("/update_product/:id", productHandler.UpdateProduct)

}
