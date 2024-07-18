package routes

import (
	"api_goshop/config"
	"api_goshop/handler"
	middlewares "api_goshop/middleware"
	"api_goshop/repositories"
	"api_goshop/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.RouterGroup) {
	userAccountRepo := repositories.NewUserRepository(config.DB)
	userAccountService := services.NewUserService(userAccountRepo)
	userAccountHandler := handler.NewUserAccount(userAccountService)
	c.POST("/register", userAccountHandler.RegisterHandler)
	c.POST("/login", userAccountHandler.LoginHandler)
	c.Use(middlewares.JWTMiddleware())
	c.GET("/user", userAccountHandler.GetDataUser)
}
