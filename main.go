package main

import (
	"api_goshop/config"
	middlewares "api_goshop/middleware"
	"api_goshop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.DBConnetcion()

	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())
	api := router.Group("/api")
	routes.UserRoutes(api)
	routes.CartRouters(api)
	router.Run()

}
