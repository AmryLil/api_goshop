package routes

import (
	"api_goshop/config"
	"api_goshop/handler"
	middlewares "api_goshop/middleware"
	"api_goshop/repositories"
	"api_goshop/services"

	"github.com/gin-gonic/gin"
)

func CartRouters(c *gin.RouterGroup) {
	cartRepo := repositories.NewCartRepository(config.DB)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)
	c.Use(middlewares.JWTMiddleware())
	c.POST("/addtocart", cartHandler.AddtoCartHandler)
	c.GET("/cart", cartHandler.ReadCartHandler)
	c.POST("/updatecart", cartHandler.UpdateCartHandler)
	c.GET("/deletecart/:id", cartHandler.DeleteCartHandler)
}
