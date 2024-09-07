package routes

import (
	"api_goshop/config"
	"api_goshop/handler"
	middlewares "api_goshop/middleware"
	"api_goshop/repositories"
	"api_goshop/services"

	"github.com/gin-gonic/gin"
)

func PaymentRouters(c *gin.RouterGroup) {
	paymentRepo := repositories.NewPaymentRepo(config.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	c.Use(middlewares.JWTMiddleware())
	c.POST("/payment", paymentHandler.CreateTransaction)
}
